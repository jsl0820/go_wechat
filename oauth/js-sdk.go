package oauth

import (
	"fmt"
	"log"
	"time"

	. "github.com/jsl0820/wechat"
)

const TICKET_URL = "/cgi-bin/ticket/getticket?type=jsapi&access_token={{TOKEN}}"

type TicketResp struct {
	ErrCode   string `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

type Ticket struct {
	expires uint
	Resp    TicketResp
}

var ticketInstance = &Ticket{expires: GetConfig().Expires}

//刷新
func (t *Ticket) Refresh() (TicketResp, error) {
	tk, err := t.Get()
	if err != nil {
		fmt.Println(err)
	}

	url := Url(TICKET_URL)
	var s TicketResp
	err = NewRequest().Get(url).JsonResp(&s)

	if err != nil {
		return s, err
	}

	return s, nil
}

//获取票据
func (t *Ticket) Get() (string, error) {
	if t.Resp == (TicketResp{}) {
		resp, err := t.Refresh()
		if err != nil {
			log.Println(err)
			return "", err
		}

		t.Resp = resp
	}

	return t.Resp.Ticket, nil
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

type JsConfigConfig struct {
	TimeStamp string `json:"timesamp"`
	NonceStr  string `json:"noncestr"`
	Signature string `json:"signature"`
}

func NewJsSdk(url string) *JsSdk {
	config := &JsConfigConfig{}
	return &JsSdk{SdkConfig: config}
}

type JsSdk struct {
	Url       string
	config    map[string]string
	SdkConfig *JsConfigConfig
}

func (sdk *JsSdk) Get() map[string]string {
	tamp := StampString()
	nonceStr := NonceStringGenerator(32)
	ticket, err := ticketInstance.Get()
	if err != nil {
		log.Println(err)
	}

	sdk.config = map[string]string{
		"url":          sdk.Url,
		"noncestr":     nonceStr,
		"timestamp":    tamp,
		"jsapi_ticket": ticket,
	}

	sign := StringSign(sdk.config)
	nature := Sha1Sign(sign)

	return map[string]string{
		"timesamp":  tamp,
		"noncestr":  nonceStr,
		"signature": nature,
	}
}

func init() {
	go ticketInstance.Clear()
}
