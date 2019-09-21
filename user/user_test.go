package user

import (
	. "github.com/jsl0820/wechat"
	"testing"
)

func init()  {
	config := Config{
		WxAppId:     "wx582ef3694f7a7546",
		WxAppSecret: "148ee9063222674ef03e4c21776e02cd",
	}

	WxConfig(config)
}

func TestCreate(t *testing.T) {
	tag := new(UserTag)
	id, err := tag.Create("dog122")
	t.Log(id)
	t.Log(err)
}

func TestList(t *testing.T) {
	tag := new(UserTag)
	resp, err := tag.List()
	t.Log(*resp)
	t.Log(err)
}

func TestDel(t *testing.T) {
	tag := new(UserTag)
	isDel := tag.Del(104)
	t.Log(isDel)
}


