package oauth

import (
	"errors"
	"log"
	"strings"

	. "github.com/jsl0820/wechat"
)

const OAUTH_URL = "/connect/oauth2/authorize?appid={{APPID}}&redirect_uri={{REDIRECT_URI}}&response_type=code&scope={{SCOPE}}&state=STATE#wechat_redirect"
const OATH_TOKEN_URL = "/sns/oauth2/access_token?appid={{APPID}}&secret={{SECRET}}&code={{CODE}}"
const USERINFO_URL = "/sns/userinfo?access_token={{TOKEN}}&openid={{OPENID}}&lang=zh_CN"
const TOKEN_EXPIRED = "/sns/auth?access_token={{TOKEN}}&openid={{OPENID}}"

// var (
// 	OAUTH_ERROR =
// )

//
func OauthUrl(url, string, t uint8) string {
	scope := "snsapi_base"
	if t == 2 {
		scope = "snsapi_userinfo"
	}

	uri := HOST + OAUTH_URL
	uri = Url(uri)
	uri = strings.Replace(uri, "{{SCOPE}}", scope, -1)
	uri = strings.Replace(uri, "{{REDIRECT_URI}}", uri, -1)
	uri = strings.Replace(uri, "{{APPID}}", GetConfig().WxAppId, -1)

	return uri
}

type SnsapiBase struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Errcode      string `json:"errcode"`
	Errmsg       string `json:"errmsg"`
}

type UserInfo struct {
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

func (o *Oauth) token() {
	url := strings.Replace(OATH_TOKEN_URL, "{{APPID}}", GetConfig().WxAppId, -1)
	url = strings.Replace(url, "{{SECRET}}", GetConfig().WxAppSecret, -1)
	url = strings.Replace(url, "{{CODE}}", o.Code, -1)

	resp := make(map[string]string)
	if err := NewRequest().Get(url).JsonResp(&resp); err != nil {
		panic(err)
	}

	if resp["errcode"] != "" {
		panic(resp["errmsg"])
	}

	o.Openid = resp["openid"]
	o.AccessToken = resp["access_token"]
	o.RefreshToken = resp["refresh_token"]
}

//验证token是否过期
func (o *Oauth) isTokenExpired() bool {
	url := HOST + TOKEN_EXPIRED
	url = strings.Replace(url, "{{OPENDID}}", o.Openid, -1)
	url = strings.Replace(url, "{{TOKEN}}", o.AccessToken, -1)

	resp := make(map[string]string)
	if err := NewRequest().Get(url).JsonResp(&resp); err != nil {
		log.Println(err)
		return false
	}

	if resp["errcode"] == "0" {
		return true
	}

	log.Println("code:", resp["code"], ", msg:"+resp["errmag"])
	return false
}

// 获取SnsapiBase
func (o *Oauth) UserInfo() (UserInfo, error) {

	//是否有效
	if !o.isTokenExpired() {
		o.token()
	}

	url := HOST + USERINFO_URL
	url = strings.Replace(url, "{{CODE}}", o.Code, -1)
	url = strings.Replace(url, "{{OPENID}}", o.Openid, -1)
	url = strings.Replace(url, "{{TOKEN}}", o.AccessToken, -1)

	var resp UserInfo
	if err := NewRequest().Get(url).JsonResp(&resp); err != nil {
		panic(err)
	}

	if resp.Errcode != "" {
		return resp, errors.New("获取SnsapiBase出错, code:" + resp.Errcode + ", msg：" + resp.Errmsg)
	}

	return resp, nil
}
