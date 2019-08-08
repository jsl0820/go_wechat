package wechat

import (
	"log"
	"testing"
	"os"
	"strings"
	"io/ioutil"
)	

var openid = "oKPxfwK493kkbIH1dBrIP-nBADBc"
//go test -v wechat_test.go  ip_list.go wechat.go http.go  token.go -run TestIp

func TestConfig(t *testing.T){
	config := Config{
		WxAppId:"wx75d0a800a00671a1",
		WxAppSecret:"de3426ea07a05887a220c91232fcc9e7",
	}
	WxConfig(config)
	log.Println(Wxconfig)
}

// file token.go 
func TestToken(t *testing.T){
	TestConfig(t)
	tk, err := token.Get()
	if err !=nil {
		log.Println(err)
	}
	
	log.Println("token", tk)
}

func TestReflash(t *testing.T){
	token := &Token{Expires: 7200}
	token.Refresh()
}

// file ip_list
//获取微信服务器列表
func TestIpList(t *testing.T){
	TestToken(t)
	ips, err := IpList()
	log.Println(ips, err)
}

// func TestKfAdd(t *testing.T){
// 	TestToken(t)
// 	kf := &KfAct{}
// 	b, err := kf.Add("user1@gh_2cd837cec3e9", "javelin")
// 	log.Println(b, err)
// }

// func TestKfGet(t *testing.T){
// 	TestToken(t)
// 	kf := &KfAct{}
// 	list, err := kf.Get()
// 	log.Println(list, err)
// }

//
// func TestUserQuery(t *testing.T){
// 	user := NewUsers()
// 	u, err :=user.Query(openid, "zh_CN")
// 	log.Println(u, err)
// }

func TestUploadImage(t *testing.T){
	config := Config{
		WxAppId:"wx582ef3694f7a7546",
		WxAppSecret:"148ee9063222674ef03e4c21776e02cd",
	}
	WxConfig(config)

	m := &Media{}
	id, err := m.UploadImg("test.jpg")
	if err != nil {
		log.Println("err", err)
	}

	log.Println(id)

}

//表单上传文件测试
// func TestFormFile(t *testing.T){
// 	req := NewRequest()
// 	req.FormFile("test", "token.go")
// 	req.FormField("username", "javelin")
// 	resp, err  := req.FormData().Post("http://httpbin.org/post").String()

// 	if err != nil {
// 		log.Println("111")
// 	}

// 	log.Println(resp)
// }

func TestSaveTo(t *testing.T){
	f := "beego_testfile"
	req := NewRequest().Get("http://httpbin.org/ip")
	err := req.SaveTo(f)
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(f)
	b, err := ioutil.ReadFile(f)
	if n := strings.Index(string(b), "origin"); n == -1 {
		t.Fatal(err)
	}
}



// func TestFileUpLoad(t *testing.T){
// 	config := Config{
// 		WxAppId:"wx582ef3694f7a7546",
// 		WxAppSecret:"148ee9063222674ef03e4c21776e02cd",
// 	}
// 	WxConfig(config)
// 	tk, _ := token.Get()
// 	log.Println(tk)
// 	url := HOST + "/cgi-bin/media/upload?access_token="+ tk + "&type=image"
// 	NewRequest().PostFile(url, "media", "test.jpg")
// }