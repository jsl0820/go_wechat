package user

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

func TestUserInfo(t *testing.T) {
	u := new(User)
	info, err := u.Info("oKPxfwCpNKAxAA01yjjWt1WJY6-k")
	t.Log(info.Remark)
	t.Log(err)
}

func TestUserList(t *testing.T) {
	u := new(User)
	list, e := u.List("oKPxfwCpNKAxAA01yjjWt1WJY6-k")
	t.Log(list)
	t.Log(e)
}

func TestUserRemark(t *testing.T) {
	u := new(User)
	isRemarked := u.Remark("oKPxfwCpNKAxAA01yjjWt1WJY6-k", "male")
	t.Log(isRemarked)
}

func TestUserBlock(t *testing.T) {
	u := new(User)
	var users = []string{
		"oKPxfwCpNKAxAA01yjjWt1WJY6-k",
		"oKPxfwK493kkbIH1dBrIP-nBADBc",
	}

	isSuccess := u.Block(users...)
	t.Log(isSuccess)
}

func TestBlackList(t *testing.T) {
	u := new(User)
	list, e := u.BlackList("oKPxfwCpNKAxAA01yjjWt1WJY6-k")
	t.Log(list)
	t.Log(e)
}

func TestBlackCancel(t *testing.T) {
	u := new(User)
	var users = []string{
		"oKPxfwCpNKAxAA01yjjWt1WJY6-k",
		"oKPxfwK493kkbIH1dBrIP-nBADBc",
	}
	isCancel := u.BlockCancel(users...)
	t.Log(isCancel)
}
