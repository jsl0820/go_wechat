package oauth

import (
	"fmt"
	"time"
)

type TicketResp struct {
	ErrCode   string `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

var ticket = &Ticket{expires:7200}

type Ticket struct {
	expires int
	Resp TicketResp 
}

//刷新
func (t *Ticket)Refresh()(TicketResp, error){
	tk, err := token.Get()
	if err != nil {
		fmt.Println(err)
	}

	url := HOST + "/cgi-bin/ticket/getticket?type=jsapi&access_token=" + tk
	var s TicketResp
	err = NewRequest().Get(url).JsonResp(&s)

	if err != nil {
		return s, err
	}

	return s, nil	
}

//获取票据
func (t *Ticket) Get() string {
	if 	t.Resp == (TicketResp{}){
		resp, err := t.Refresh()
		if err != nil {
			fmt.Println(err)
			return ""
		}

		t.Resp = resp
		return  t.Resp.Ticket 
	} 	

	return t.Resp.Ticket
}


//定期清理
func (t *Ticket) Clear() {
	dur := time.Duration(t.expires) * time.Second
	for {
		<-time.After(dur)
		if t.Resp != (TicketResp{}) {
			t.Resp = TicketResp{}
		}
	}
}


type JsConfigConfig struct{
	TimeStamp string `json:"timesamp"`	
	NonceStr string  `json:"noncestr"`
	Signature string `json:"signature"`
}

func NewJsSdk(url string) *JsSdk{
	config := &JsConfigConfig{}
	return &JsSdk{SdkConfig:config}
}

type JsSdk struct {
	Url       string
	config map[string]string
	SdkConfig *JsConfigConfig
}

func (j *JsSdk)Get() map[string]string {

	timeSamp := StampString()
	nonceStr := NonceStringGenerator(32)

	j.config = map[string]string{
		"url":       j.Url,
		"noncestr":  nonceStr,
		"timestamp": timeSamp,
		"jsapi_ticket" : ticket.Get(),
	}

	stringSign := StringSign(j.config)
	signature := Sha1Sign(stringSign)

	return map[string]string{
		"timesamp" : timeSamp,
		"noncestr" : nonceStr,
		"signature" : signature,
	}
}

func Ticket()string{
	
}

func init() {
	go ticket.Clear()		
}