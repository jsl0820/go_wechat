package oauth

import (
	"log"
	"testing"

	. "github.com/jsl0820/wechat"
)

func TestConfig(t *testing.T) {
	config := Config{
		WxAppId:     "wx582ef3694f7a7546",
		WxAppSecret: "148ee9063222674ef03e4c21776e02cd",
	}
	WxConfig(config)
	t.Log(Wxconfig)
}

func TestToken(t *testing.T) {
	TestConfig(t)
	token, err := GetToken()
	if err != nil {
		t.Log(err)
	}

	t.Log("从服务器获得到的", token)

	if token == "" {
		t.Error("error")
	}
}

func TestUrl(t *testing.T) {
	TestConfig(t)
	url := Url(CALLBACK_IPS)
	log.Println("组装后的url", url)
}

func TestIp(t *testing.T) {
	TestConfig(t)
	if list, err := IP().List(); err != nil {
		t.Log(err)
	} else {
		t.Log(list)
	}
}

func TestHasIp(t *testing.T) {
	TestConfig(t)
	ip := "223.166.222.100"
	has := IP().Has(ip)
	t.Log("IP:", has)
}

func TestSdkConfig(t *testing.T) {
	TestConfig(t)
	url := "www.test.com"
	sign := SdkConfig(url)
	t.Log(sign)
}
