package menu

import (
	"encoding/json"
	"log"

	. "github.com/jsl0820/wechat"
	"github.com/jsl0820/wechat/oauth"
)

const MENU_DEL = "/cgi-bin/menu/delete?access_token={{TOKEN}}"
const MENU_CREATE = "/cgi-bin/menu/create?access_token={{TOKEN}}"
const MENU_CURRNT = "/cgi-bin/get_current_selfmenu_info?access_token={{TOKEN}}"

type Item map[string]interface{}

type MenueResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Menuid  uint
}

//查询
func Current() (string, error) {
	url := oauth.Url(MENU_CURRNT)
	resp, err := NewRequest().Get(url).String()
	if err != nil {
		return "", err
	}

	return resp, nil
}

//创建
func Create(menu Item) bool {
	body, err := json.Marshal(menu)
	if err != nil {
		log.Println(err)
	}

	url := oauth.Url(MENU_CREATE)
	var resp MenueResp
	request := NewRequest().Body(body)
	request.ContentType("application/json")
	if err := request.Post(url).JsonResp(&resp); err != nil {
		log.Println(err)
		return false
	}

	if resp.ErrCode != 0 {
		log.Println(resp.ErrMsg)
		return false
	}

	return true
}

//删除全部菜单
func Del() bool {
	url := oauth.Url(MENU_DEL)

	var resp MenueResp
	err := NewRequest().Get(url).JsonResp(&resp)
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
