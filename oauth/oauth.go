package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	. "github.com/jsl0820/wechat"
)

const SNSAPIAUTH_URL = "/sns/auth?access_token={{TOKEN}}&openid={{OPENID}}"
const SNSAPIBASE_URL = "/sns/oauth2/access_token?appid={{APPID}}&secret={{SECRET}}&code={{CODE}}"
const SNSAPIINFO_URL = "/sns/userinfo?access_token={{TOKEN}}&openid={{OPENID}}&lang=zh_CN"

type SnsapiBase struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Errcode      string `json:"errcode"`
	Errmsg       string `json:"errmsg"`
}

type SnsapiUserInfo struct {
	Openid     string        `json:"openid"`
	Nickname   string        `json:"nickname"`
	Sex        int           `json:"sex"`
	Province   string        `json:"province"`
	City       string        `json:"city"`
	Country    string        `json:"country"`
	Headimgurl string        `json:"headimgurl"`
	Privilege  []interface{} `json:"privilege"`
	Unionid    string        `json:"unionid"`
	Errcode    string        `json:"errcode"`
	Errmsg     string        `json:"errmsg"`
}

func New(code string) *Oauth {
	return &Oauth{Code: code}
}

type Oauth struct {
	Code         string
	Openid       string
	AccessToken  string
	RefreshToken string
}

// 获取SnsapiBase
func (o *Oauth) SnsapiBase() (SnsapiBase, error) {
	var err error

	url := Url(SNSAPIBASE_URL)
	url = strings.Replace(SNSAPIBASE_URL, "{{CODE}}", o.Code, -1)
	url = strings.Replace(SNSAPIBASE_URL, "{{APPID}}", GetConfig().WxAppId, -1)
	url = strings.Replace(SNSAPIBASE_URL, "{{SECRET}}", GetConfig().WxAppSecret, -1)

	var s SnsapiBase
	json.Unmarshal(body, &s)

	fmt.Println(s.Errcode)

	if s.Errcode != "" {
		err = errors.New("获取SnsapiBase出错, code:" + s.Errcode + ", msg：" + s.Errmsg)
	} else {
		o.AccessToken = s.AccessToken
		o.Openid = s.Openid
	}

	fmt.Println(err)
	return s, err
}

//获取UserInfo
func (o *Oauth) SnsapiUserInfo() (SnsapiUserInfo, error) {
	var err error
	// o.RefreshTokenAction()
	fmt.Println("access_token:", o.AccessToken)
	var info SnsapiUserInfo
	json.Unmarshal(body, &info)

	return info, err
}

// 验证AccessToken是否有效
func (o *Oauth) RefreshTokenAction() bool {

	isExpires := false

	url := "https://api.weixin.qq.com/sns/auth"
	url += "?access_token=" + o.AccessToken + "&openid=" + o.Openid

	type AccessToken struct {
		Errcode string `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var t AccessToken
	json.Unmarshal(body, &t)

	if t.Errcode != "0" || t.Errmsg != "ok" {
		isExpires = true
		url := "https://api.weixin.qq.com/sns/oauth2/refresh_token?"
		url += "appid=" + o.Openid + "&grant_type=refresh_token&refresh_token=" + o.RefreshToken
		resp, err := http.Get(url)

		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
		}

		var s SnsapiBase
		json.Unmarshal(body, &s)

		if s.Errcode != "" {
			err := errors.New("获取SnsapiBase出错, code:" + s.Errcode + ", msg：" + s.Errmsg)
			fmt.Println(err)
		} else {
			o.AccessToken = s.AccessToken
			o.RefreshToken = s.RefreshToken
		}
	}
	return isExpires
}
