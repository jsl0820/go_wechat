package user

import (
	"encoding/json"
	"errors"
	"github.com/jsl0820/wechat/oauth"
	"log"
	"strings"

	. "github.com/jsl0820/wechat"
)

const USER_LIST_URL = "/cgi-bin/user/get?access_token={{TOKEN}}&next_openid=NEXT_OPENID"
const USER_REMARK_URL = "/cgi-bin/user/info/updateremark?access_token={{TOKEN}}"
const USER_INFO_URL = "/cgi-bin/user/info?access_token={{TOKEN}}&openid=OPENID&lang=zh_CN"
const USER_PACKAGE_URL = "/cgi-bin/user/info/batchget?access_token={{TOKEN}}"
const USER_BLACK_LIST = "/cgi-bin/tags/members/getblacklist?access_token={{TOKEN}}"
const USER_BLACK = "/cgi-bin/tags/members/batchblacklist?access_token={{TOKEN}}"
const USER_BLACK_CANCEL = "/cgi-bin/tags/members/batchunblacklist?access_token={{TOKEN}}"

type UserResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type Information struct {
	Subscribe      int      `json:"subscribe"`
	OpenId         string   `json:"openid"`
	NickName       string   `json:"nickname"`
	Sex            int      `json:"sex"`
	Language       string   `json:"language"`
	City           string   `json:"city"`
	Province       string   `json:"province"`
	Country        string   `json:"country"`
	HeadImgUrl     string   `json:"headimgurl"`
	SubscribeTime  int64    `json:"subscribe_time"`
	Unionid        string   `json:"unionid"`
	Remark         string   `json:"remark"`
	Groupid        int      `json:"groupid"`
	TagidList      []string `json:"tagid_list"`
	SubscribeScene string   `json:"subscribe_scene"`
	QrScenetr      string   `json:"qr_scene_str"`
	ErrCode        int      `json:"errcode"`
	ErrMsg         string   `json:"errmsg"`
}

type User struct{}

type ListResp struct {
	Total      uint64
	Count      uint64
	Data       map[string][]string
	NextOpenid string
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

//关注着列表
func (u *User) List(nestId string) (*ListResp, error) {
	url := oauth.Url(USER_LIST_URL)
	url = strings.Replace(url, "NEXT_OPENID", nestId, -1)
	var resp ListResp
	err := NewRequest().Get(url).JsonResp(&resp)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.ErrCode != 0 {
		return nil, errors.New(resp.ErrMsg)
	}

	return &resp, nil
}

//黑名单
func (u *User) BlackList(id string) (*ListResp, error) {
	url := oauth.Url(USER_BLACK_LIST)
	var resp ListResp
	body := strings.Replace(`{"begin_openid":"OPENID"}`, "OPENID", id, -1)
	request := NewRequest().Body(body)
	request.ContentType("application/json")
	err := request.Post(url).JsonResp(&resp)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

//设置用户备注名
func (u *User) Remark(id, mark string) bool {
	data := make(map[string]string)
	data["openid"] = id
	data["remark"] = mark
	url := oauth.Url(USER_REMARK_URL)
	body, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return false
	}

	var resp UserResp
	request := NewRequest().Body(body)
	request.ContentType("application/json")
	err = request.Post(url).JsonResp(&resp)
	if err != nil {
		log.Println(err)
		return false
	}

	if resp.ErrCode != 0 {
		log.Println(resp.ErrMsg)
		return false
	}

	return true
}

//会员信息
func (u *User) Info(id string) (*Information, error) {
	url := oauth.Url(USER_INFO_URL)
	url = strings.Replace(url, "OPENID", id, -1)
	log.Println("请求", url)

	var resp Information
	err := NewRequest().Get(url).JsonResp(&resp)

	if err != nil {
		return nil, err
	}

	if resp.ErrCode != 0 {
		return nil, errors.New(resp.ErrMsg)
	}
	return &resp, nil
}

//批量获取用户信息
func (u *User) InfoList(ids ...string) (*map[string][]Information, error) {
	data := make(map[string]interface{})
	var users []interface{}
	for i := 0; i < len(ids); i++ {
		users = append(users, map[string]string{
			"lang":   "zh_CN",
			"openid": ids[i],
		})
	}

	data["user_list"] = users
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	url := oauth.Url(USER_PACKAGE_URL)
	var resp *map[string][]Information
	request := NewRequest().Body(body)
	request.ContentType("application/json")
	err = request.Post(url).JsonResp(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//拉黑
func (u *User) Block(ids ...string) bool {
	if len(ids) > 20 {
		log.Println("一次只能拉黑20个")
		return false
	}

	url := oauth.Url(USER_BLACK)
	data := make(map[string][]string)
	data["openid_list"] = ids
	body, e := json.Marshal(data)
	if e != nil {
		log.Println(e)
		return false
	}

	var resp UserResp
	request := NewRequest().Body(body)
	request.ContentType("application/json")
	err := request.Post(url).JsonResp(&resp)
	if err != nil {
		log.Println(err)
		return false
	}

	if resp.ErrCode != 0 {
		log.Println(resp.ErrMsg)
		return false
	}

	return true
}

//拉黑
func (u *User) BlockCancel(ids ...string) bool {
	if len(ids) > 20 {
		log.Println("一次只能拉黑20个")
		return false
	}

	url := oauth.Url(USER_BLACK_CANCEL)
	data := make(map[string][]string)
	data["openid_list"] = ids
	body, e := json.Marshal(data)
	if e != nil {
		log.Println(e)
		return false
	}

	var resp UserResp
	request := NewRequest().Body(body)
	request.ContentType("application/json")
	err := request.Post(url).JsonResp(&resp)
	if err != nil {
		log.Println(err)
		return false
	}

	if resp.ErrCode != 0 {
		log.Println(resp.ErrMsg)
		return false
	}

	return true
}
