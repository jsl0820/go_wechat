package wechat

import (
	"net/http"
	"fmt"	
	"io/ioutil"
	"encoding/json"
)

type JsapiTicket struct{
	ErrCode 	string	`json:"errcode"`
	ErrMsg 		string	`json:"errmsg"`
	Ticket 		string	`json:"ticket"`
	ExpiresIn 	int 	`json:"expires_in"`	
}

type JsSdk struct {
	Token  string
	Jt JsapiTicket
} 

func NewTicket(token string) *JsSdk{
	return &JsSdk{
		Token : token,
	}
}

func (js *JsSdk)Get() JsapiTicket {
	url := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?"
	url += "type=jsapi&access_token=" + js.Token
	resp, err := http.Get(url)	
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)	
	if err != nil {
		fmt.Println(err)		
	}
	var s JsapiTicket
	json.Unmarshal(body, &s)
	return s	
}

