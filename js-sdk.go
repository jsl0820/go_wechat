package wechat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JsapiTicket struct {
	ErrCode   string `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

type Ticket struct {
	Noncestr string `json:"noncestr"`
	Jt string `json:"jsapi_ticket"`
	TimeStamp int64 `json:"timestamp"`
	Url string `json:"url"`  
}


func Ticket() {
	t := 
	url := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?"
	url += "type=jsapi&access_token=" + token

	var s JsapiTicket
	err := NewRequest(&s).Get(url)

	if err != nil {
		return nil, error
	}

	return s, nil
}
