package menu

import (
	"encoding/json"
	"log"

	. "github.com/jsl0820/wechat"
	"github.com/jsl0820/wechat/oauth"
)

const COND_MENU_ADD = "/cgi-bin/menu/addconditional?access_token={{TOKEN}}"
const COND_MENU_DEL = "/cgi-bin/menu/delconditional?access_token={{TOKEN}}"
const COND_MENU_TEST = "/cgi-bin/menu/trymatch?access_token={{TOKEN}}"

func CreateCondMenu(menu Item) (uint, bool) {
	b, e := json.Marshal(menu)
	if e != nil {
		log.Println(e)
		return 0, false
	}

	url := oauth.Url(COND_MENU_ADD)
	var resp MenueResp
	request := NewRequest().Body(b)
	request.ContentType("application/json")
	if err := request.Post(url).JsonResp(&resp); err != nil {
		log.Println(err)
		return 0, false
	}

	if resp.ErrCode != 0 {
		log.Println(resp.ErrMsg)
		return 0, false
	}

	return resp.Menuid, true
}

//测试个性菜单
func CondMenu(user map[string]string) (string, error) {
	b, e := json.Marshal(user)
	if e != nil {
		log.Println(e)
		return "", e
	}

	url := oauth.Url(COND_MENU_TEST)
	request := NewRequest().Body(b)
	request.ContentType("application/json")
	resp, err := request.Post(url).String()
	if err != nil {
		return "", err
	}

	return resp, nil
}
