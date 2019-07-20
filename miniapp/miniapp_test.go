package miniapp

import (
	// "encoding/json"
	"fmt"
	"testing"
)

type value map[string]string

func Test_GetAccessToken(t *testing.T) {
	info := AppInfo{
		AppId:     "wx7fcb8e375b63648e",
		AppSecret: "eeae87d6125c8c2b1eaaea9962133fe0",
	}
	token := GetAccessToken(info)
	fmt.Println(token.ExpiresIn)
}

func Test_PushTemplateMsg(t *testing.T) {
	info := AppInfo{
		AppId:     "wx7fcb8e375b63648e",
		AppSecret: "eeae87d6125c8c2b1eaaea9962133fe0",
	}

	data := make(map[string]interface{})
	data["keyword1"] = value{"value": "11"}
	data["keyword2"] = value{"value": "内容"}
	data["keyword3"] = value{"value": "2018-12-15"}

	config := TempMsgConfig{
		FromId:     "98e52625ff322995757696b4b00ca1e8",
		TemplateId: "4Ct65ry8LW-9iUBxIhTqTSk899xo-ymsHPx85xhJhVA",
		Touser:     "onUg-5bbAmk_9C3iF-6f-RYVdZuE",
		Data:       data,
	}
	res := PushTemplateMsg(info, config)
	fmt.Println(res.ErrMsg)
}

func Test_GetMpUserId(t *testing.T) {
	info := AppInfo{
		AppId:     "wx7fcb8e375b63648e",
		AppSecret: "eeae87d6125c8c2b1eaaea9962133fe0",
	}
	code := "011hB3GS0ZLM4Y1uk7GS02NfGS0hB3GG"
	res := GetMpUserId(code, info)
	fmt.Println("aa", res)
}
