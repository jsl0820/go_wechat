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

func Ticket(token string) {

	url := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?"
	url += "type=jsapi&access_token=" + token

	var s JsapiTicket
	err := NewRequest(&s).Get(url)

	if err != nil {
		return nil, error
	}

	return s, nil
}
