package Qrcode

import (
	. "github.com/jsl0820/wechat"
	"testing"
)

func init() {
	config := Config{
		WxAppId:     "wx582ef3694f7a7546",
		WxAppSecret: "148ee9063222674ef03e4c21776e02cd",
	}

	WxConfig(config)
}

func TestTicket(t *testing.T) {
	param := make(map[string]interface{})
	scene := make(map[string]interface{})
	scene["scene"] = map[string]string{
		"scene_str": "test",
	}
	param["expire_seconds"] = 604800
	param["action_name"] = "QR_STR_SCENE"
	param["action_info"] = scene

	ticket, err := New(param).Create()
	t.Log(err)
	t.Log(ticket)
}

func TestUrl(t *testing.T) {
	param := make(map[string]interface{})
	scene := make(map[string]interface{})
	scene["scene"] = map[string]string{
		"scene_str": "test222",
	}
	param["expire_seconds"] = 604800
	param["action_name"] = "QR_STR_SCENE"
	param["action_info"] = scene

	ur := New(param).Url()
	t.Log(ur)
}

func TestShortUrl(t *testing.T) {
	l, e := ShortUrl("https://jsl0820.github.io/")
	t.Log(l)
	t.Log(e)
}
