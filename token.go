package wechat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

const (
	TOKEN_ERROR_1 = "系统繁忙，此时请开发者稍候再试"
	TOKEN_ERROR_2 = "AppSecret错误或者AppSecret不属于这个公众号，请开发者确认AppSecret的正确性"
	TOKEN_ERROR_3 = "请确保grant_type字段值为client_credential"
	TOKEN_ERROR_4 = "调用接口的IP地址不在白名单中，请在接口IP白名单中进行设置。（小程序及小游戏调用不要求IP地址在白名单内"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

type Token struct {
	At AccessToken
}

var token Token

//token 定期清理时间 
var TokenGcTime = 7200


//刷新token
func (t *Token)Reflash() error {

	url := "https://api.weixin.qq.com/cgi-bin/token?"
	url += "grant_type=client_credential&appid=" + WxAppId + "&secret=" + WxAppSecret

	err := NewRequest(&t.At).Get(url)
	if err != nil {
		return err 
	}
	return nil
}

//获取token
func (t *Token) Get() (string, error ){
	if t.AccessToken == (AccessToken{}) {
		err := t.Reflash()
		if err != nil {
			return nil, err
		}

		switch t.At.Errcode{
		case -1 :
			return errors.New(TOKEN_ERROR_1)
		case 40001:
			return errors.New(TOKEN_ERROR_2)	
		case 40002:
			return 	errors.New(TOKEN_ERROR_3)
		case 40164:	
			return errors.New(TOKEN_ERROR_4)
		default :
			return t.At.AccessToken
		}
	} 

	return t.At.AccessToken
}

//定期清空 时间间隔为 TokenGcTime
func (t. Token) clear(){
	duration := time.Duration(TokenGcTime) * time.Second
	for {
		<-time.After(duration)
		if t.At != (AccessToken{}) {
			t.At = AccessToken{} 
		}
	}		
}


func init(){
	token = &Token{}
	go token.clear()
}

