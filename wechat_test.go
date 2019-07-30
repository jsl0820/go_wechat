package wechat

import (
	"fmt"
	"testing"
)	

var openid = "oKPxfwK493kkbIH1dBrIP-nBADBc"


//测试单个方法 go test -v -test.run 方法名称


// file token.go 
func TestToken(t *testing.T){
	tk, _ := token.Get()
	fmt.Println("token", tk)
}

func TestReflash(t *testing.T){
	token := &Token{Expires: 7200}
	token.Refresh()
}


// file ip_list
//获取微信服务器列表
func TestIpList(t *testing.T){
	ips, err := IpList()
	fmt.Println(ips, err)
}

// file Kf

func TestKfAdd(t *testing.T){
	kf := &KfAct{}
	b, err := kf.Add("test3@test.com", "客服222")
	fmt.Println(b, err)
}

//
func TestUserQuery(t *testing.T){
	user := NewUsers()
	u, err :=user.Query(openid, "zh_CN")
	fmt.Println(*u, err)
}


func TestUserQueryAll(t *testing.T){
	
}