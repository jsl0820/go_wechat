package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func NewOauth() *Oauth {
	return new(Oauth)
}

type Oauth struct {
	Code         string
	Openid       string
	AccessToken  string
	RefreshToken string
}

// 获取SnsapiBase
func (o *Oauth) GetSnsapiBase() (SnsapiBase, error) {
	fmt.Println(Appid, AppSecret)
	var err error
	url := HOST + "/sns/oauth2/access_token?"
	url += "appid=" + Appid + "&secret=" + AppSecret + "&code=" + o.Code
	url += "&grant_type=authorization_code"

	resp, err := http.Get(url)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

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
func (o *Oauth) GetSnsapiUserInfo() (SnsapiUserInfo, error) {
	var err error
	// o.RefreshTokenAction()

	url := "https://api.weixin.qq.com/sns/userinfo?access_token="
	url += o.AccessToken + "&openid=" + o.Openid + "&lang=zh_CN"

	fmt.Println("access_token:", o.AccessToken)

	resp, err := http.Get(url)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

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
