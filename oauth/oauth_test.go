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

func TestUrl(t *testing.T) {
	TestConfig(t)
	url := Url(CALLBACK_IPS)
	log.Println("组装后的url", url)
}

func TestIp(t *testing.T) {
	if list, err := IP().List(); err != nil {
		log.Println(err)
	} else {
		log.Println(list)
	}
}

func TestHasIp(t *testing.T) {
	// TestConfig(t)
	ip := "223.166.222.100"
	has := IP().Has(ip)
	t.Log("IP:", has)
}

func TestSdkConfig(t *testing.T) {
	url := "www.test.com"
	sign := SdkConfig(url)
	log.Println(sign)
}
