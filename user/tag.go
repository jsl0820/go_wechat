package user

import (
	"log"
	"strconv"
	"strings"

	. "github.com/jsl0820/wechat"
	"github.com/jsl0820/wechat/oauth"
)

const USER_TAG_DEL = "/cgi-bin/tags/delete?access_token={{TOKEN}}"
const USER_TAG_ALL = "/cgi-bin/tags/get?access_token={{TOKEN}}"
const USER_TAG_ADD = "/cgi-bin/tags/create?access_token={{TOKEN}}"

type Tag struct {
	Id uint64
	Name string
	Count uint64
}

type TagResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Tag Tag
}

type UserTag struct {}

//创建标签
func(u *UserTag) Create(tag string) (uint64, error) {
	url := oauth.Url(USER_TAG_ADD)
	body := strings.Replace(`{"tag":{"name":"TAG"}}`, "TAG", tag, -1)

	var resp TagResp
	request := NewRequest().Body(body)
	request.ContentType("application/json")
	err := request.Post(url).JsonResp(&resp)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	if resp.ErrCode != 0 {
		log.Println(resp.ErrMsg)
		return 0, err
	}

	log.Println(resp)
	return resp.Tag.Id, nil
}

//获取标签
func (u *UserTag)List() (*map[string][]Tag, error) {
	url := oauth.Url(USER_TAG_ALL)
	var resp = new(map[string][]Tag)
	err := NewRequest().Get(url).JsonResp(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 删除
func (u *UserTag)Del(id uint64) bool {
	idString := strconv.FormatUint(id, 10)
	url := oauth.Url(USER_TAG_DEL)
	body := strings.Replace(`{"tag":{"id":ID}}`, "ID", idString, -1)
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
