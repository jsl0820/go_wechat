package wechat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

func NewToken() (AccessToken, err) {

	url := "https://api.weixin.qq.com/cgi-bin/token?"
	url += "grant_type=client_credential&appid=" + WxAppId + "&secret=" + WxAppSecret

	var at AccessToken

	err := NewRequest(&at).Get(url)
	if err != nil {
		return nil, err
	}

	return at, nil
}
