package oauth

import (
	"errors"
	"time"

	. "github.com/jsl0820/wechat"
)

const TICKET_URL = "/cgi-bin/ticket/getticket?type=jsapi&access_token={{TOKEN}}"

var ticketInstance = &Ticket{Expires: GetConfig().Expires}

type Ticket struct {
	Expires uint
	Ticket  string
}

//刷新票据
func (ti *Ticket) ticketRefresh() {
	url := Url(TICKET_URL)
	type Resp struct {
		Errcode   int
		Errmsg    string
		Ticket    string
		ExpiresIn uint
	}

	var resp Resp
	if err := NewRequest().Get(url).JsonResp(&resp); err != nil {
		panic(err)
	}

	if resp.Errcode != 0 {
		panic(errors.New("errmsg:" + resp.Errmsg))
	}

	ti.Ticket = resp.Ticket
}

//定期清理
func (ti *Ticket) Clear() {
	d := time.Duration(ti.Expires) * time.Second
	for {
		<-time.After(d)
		if ti.Ticket != "" {
			ti.Ticket = ""
		}
	}
}

//获取
func (ti *Ticket) GetTicket() string {
	if ti.Ticket == "" {
		ti.ticketRefresh()
	}

	return ti.Ticket
}

//js-sdk配置
func SdkConfig(url string) map[string]string {
	m := make(map[string]string)
	sign := make(map[string]string)

	m["url"] = url
	m["timestamp"] = StampString()
	m["noncestr"] = NonceStringGenerator(32)
	m["jsapi_ticket"] = ticketInstance.GetTicket()

	sign["timestamp"] = m["timestamp"]
	sign["noncestr"] = m["noncestr"]
	sign["signature"] = Sha1Sign(StringSign(m))

	return sign
}

func init() {
	go ticketInstance.Clear()
}
