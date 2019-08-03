package wechat

import (
	"log"
	"testing"
	// "os"
	"strings"
	"io/ioutil"
)	

var openid = "oKPxfwK493kkbIH1dBrIP-nBADBc"


//测试单个方法 go test -v -test.run 方法名称


// file token.go 
func TestToken(t *testing.T){
	tk, _ := token.Get()
	log.Println("token", tk)
}

func TestReflash(t *testing.T){
	token := &Token{Expires: 7200}
	token.Refresh()
}


// file ip_list
//获取微信服务器列表
func TestIpList(t *testing.T){
	ips, err := IpList()
	log.Println(ips, err)
}

// file Kf

func TestKfAdd(t *testing.T){
	kf := &KfAct{}
	b, err := kf.Add("test3@test.com", "客服222")
	log.Println(b, err)
}

//
func TestUserQuery(t *testing.T){
	user := NewUsers()
	u, err :=user.Query(openid, "zh_CN")
	log.Println(u, err)
}


func TestUserQueryAll(t *testing.T){
	
}

//表单上传文件测试
func TestFormFile(t *testing.T){
	req := NewRequest()
	req.FormFile("test", "token.go")
	req.Param("username", "javelin")
	resp, err  := req.Form().Post("http://httpbin.org/post").String()

	if err != nil {
		log.Println("111")
	}

	log.Println(resp)
}

//
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