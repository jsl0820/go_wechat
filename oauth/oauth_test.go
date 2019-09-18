package oauth

import (
	"log"
	"testing"

	. "github.com/jsl0820/wechat"
)

func TestConfig(t *testing.T) {
	config := Config{
		WxAppId:     "wx75d0a800a00671a1",
		WxAppSecret: "de3426ea07a05887a220c91232fcc9e7",
	}
	WxConfig(config)
	log.Println(Wxconfig)
}

func TestToken(t *testing.T) {
	// TestConfig(t)
	token, err := GetToken()
	if err != nil {
		log.Println(err)
	}

	log.Println("从服务器获得到的", token)
}
