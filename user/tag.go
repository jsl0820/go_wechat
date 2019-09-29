package user

import (
	"log"
	"strings"

	. "github.com/jsl0820/wechat"
	"github.com/jsl0820/wechat/oauth"
)

const USER_TAG_DEL = "/cgi-bin/tags/delete?access_token={{TOKEN}}"
const USER_TAG_ALL = "/cgi-bin/tags/get?access_token={{TOKEN}}"
const USER_TAG_ADD = "/cgi-bin/tags/create?access_token={{TOKEN}}"

type TagResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Tag     Tag
}

type Tag struct {
	Id    uint
	Name  string
	Count uint
}

type UserTag struct {
}

//创建标签
func (u *UserTag) Create(tag string) bool {
	url := oauth.Url(USER_TAG_ADD)
	body := strings.Replace(`{"tag":{"name":TAG}}`, "TAG", tag, -1)
	var resp TagResp
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

//获取标签
func List() (*map[string][]Tag, error) {
	url := oauth.Url(USER_TAG_ALL)
	var resp *map[string][]Tag
	err := NewRequest().Get(url).JsonResp(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 删除
func Tag(id string) bool {
	url := oauth.Url(USER_TAG_DEL)
	body := strings.Replace(`{"tag":{"id:ID}}`, "ID", id, -1)

	var resp TagResp
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
