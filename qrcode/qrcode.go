package Qrcode

import (
	"encoding/json"
	"errors"
	. "github.com/jsl0820/wechat"
	"github.com/jsl0820/wechat/oauth"
	"log"
	url2 "net/url"
	"strings"
)

const SHORT_URL = "/cgi-bin/shorturl?access_token={{TOKEN}}"
const QRCODE_INFO = "/cgi-bin/showqrcode?ticket="
const QRCODE_TICKET = "/cgi-bin/qrcode/create?access_token={{TOKEN}}"

type CodeInfo struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds uint64 `json:"expire_seconds"`
	Url           string `json:"url"`
	ErrCode       int    `json:"errcode"`
	ErrMsg        string `json:"errmsg"`
	ShortUrl      string `json:"short_url"`
}

func New(p interface{}) *Qrcode {
	b, e := json.Marshal(p)
	if e != nil {
		panic(e)
	}

	return &Qrcode{CodeJson: b}
}

type Qrcode struct {
	CodeJson []byte
	Info     CodeInfo
}

//获取ticket
func (q *Qrcode) Create() (*CodeInfo, error) {
	url := oauth.Url(QRCODE_TICKET)
	var resp CodeInfo
	request := NewRequest().Body(q.CodeJson)
	request.ContentType("application/json")
	err := request.Post(url).JsonResp(&resp)
	if err != nil {
		return nil, err
	}

	q.Info = resp
	return &resp, nil
}

//二维码路径
func (q *Qrcode) Url() string {
	q.Create()
	v := url2.Values{}
	v.Add("ticket", q.Info.Ticket)
	value, err := url2.ParseQuery(v.Encode())
	if err != nil {
		panic(err)
	}

	log.Printf("%#v", value)
	return HOST + QRCODE_INFO + value["ticket"][0]
}

//保存到
func (q *Qrcode) ToFile(p string) bool {
	if err := NewRequest().Get(q.Url()).SaveTo("./" + p); err != nil {
		return false
	}

	return true
}

type ShortUrlResp struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	ShortUrl string `json:"short_url"`
}

//长连接转短链接
func ShortUrl(url string) (string, error) {
	url1 := oauth.Url(SHORT_URL)
	body := `{"action":"long2short", "long_url":"URL"}`
	body = strings.Replace(body, "URL", url, -1)

	var resp ShortUrlResp
	request := NewRequest().Body(body)
	request.ContentType("application/json")
	err := request.Post(url1).JsonResp(&resp)
	if err != nil {
		return "", err
	}

	if resp.ErrCode != 0 {
		return "", errors.New(resp.ErrMsg)
	}

	return resp.ShortUrl, nil
}
