package oauth

import (
	"errors"
	"strings"
	"time"

	. "github.com/jsl0820/wechat"
)

const TOKEN_URL = "/cgi-bin/token?grant_type=client_credential&appid={{APPID}}&secret={{APP_SECRITE}}"

var (
	TOKEN_ERROR_1 = "系统繁忙，此时请开发者稍候再试"
	TOKEN_ERROR_2 = "AppSecret错误或者AppSecret不属于这个公众号，请开发者确认AppSecret的正确性"
	TOKEN_ERROR_3 = "请确保grant_type字段值为client_credential"
	TOKEN_ERROR_4 = "调用接口的IP地址不在白名单中，请在接口IP白名单中进行设置。小程序及小游戏调用不要求IP地址在白名单内"
)

type TokenResp struct {
	Errmsg      string `json:"errmsg"`
	Errcode     int    `json:"errcode"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type Token struct {
	Expires     uint
	AccessToken TokenResp
}

var TokenInstance = &Token{Expires: GetConfig().Expires}

//刷新token
func (t *Token) Refresh() error {
	url := HOST + TOKEN_URL
	url = strings.Replace(url, "{{APPID}}", GetConfig().WxAppId, -1)
	url = strings.Replace(url, "{{APP_SECRITE}}", GetConfig().WxAppSecret, -1)
	if err := NewRequest().Get(url).JsonResp(&t.AccessToken); err != nil {
		panic(err)
		return err
	}

	return nil
}

//获取token
func (t *Token) Get() (string, error) {
	if t.AccessToken == (TokenResp{}) {
		if err := t.Refresh(); err != nil {
			return "", err
		}

		switch t.AccessToken.Errcode {
		case -1:
			return "", errors.New(TOKEN_ERROR_1)
		case 40001:
			return "", errors.New(TOKEN_ERROR_2)
		case 40002:
			return "", errors.New(TOKEN_ERROR_3)
		case 40164:
			return "", errors.New(TOKEN_ERROR_4)
		default:
			return t.AccessToken.AccessToken, nil
		}
	}

	return t.AccessToken.AccessToken, nil
}

//定期清空 时间间隔为 TokenGcTime
func (t *Token) clean() {
	duration := time.Duration(t.Expires) * time.Second
	for {
		<-time.After(duration)
		if t.AccessToken != (TokenResp{}) {
			t.AccessToken = TokenResp{}
		}
	}
}

//构建请求
func Url(url string) string {
	token, err := TokenInstance.Get()
	if err != nil {
		panic(err)
	}

	return strings.Replace(url, "{{TOKEN}}", token, -1)
}

func init() {
	go TokenInstance.clean()
}
