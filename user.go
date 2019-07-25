package wechat

import (
	"fmt"
	"encoding/json"
)

type UserResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type User struct {
	Subscribe      int      `json:"subscribe"`
	OpenId         string   `json:"openid"`
	NickName       string   `json:"nickname"`
	Sex            int      `json:"sex`
	Language       string   `json:"language"`
	City           string   `json:"city"`
	Province       string   `json:"province"`
	Country        string   `json:"country"`
	HeadImgUrl     string   `json:"headimgurl"`
	SubscribeTime  string   `json:"subscribe_time"`
	Unionid        string   `json:"unionid"`
	Remark         string   `json:"remark"`
	Groupid        []string `json:"groupid"`
	TagidList      string   `json:"tagid_list"`
	SubscribeScene string   `json:"subscribe_scene"`
	QrScenetr      string   `json:"qr_scene_str"`
}

type UserList struct {
	List []map[string]string `json:"user_list"`
}

type UserListResp struct {
	List []User `json:"user_info_list"` 
}

func NewUsers() *Users{
	t, _ := token.Get()
	return &Users{token:t}
}

//用户
type Users struct{
	token string
}

//用户备注
func (u *Users)Remark(openid, remark string) bool {
	url := HOST + "/cgi-bin/user/info/updateremark?access_token=" + u.token
	body := `{
		openid:{{.openid}},
		remark:{{.remark}}
	}`
	
	var resp UserResp 
	err := NewRequest().Body(body).JsonResp(&resp)
	if err != nil {
		return false
	}	

	if resp.ErrCode != 0{
		return false
	} 

	return true
}

//获取用户信息
func (u *Users)Query(openid string, lang string) (*User, bool) {
	url := HOST + "/cgi-bin/user/info?access_token="
	url += u.token + "&openid=" + openid + "&lang=" + lang

	var user User
	err := NewRequest().Get(url).JsonResp(&user)
	if err != nil {
		fmt.Println(err)
		return nil, false 
	}

	return &user, true
}

//批量获取用户信息
func (u *Users)QueryAll(ids []map[string]string) (*[]User, error) {
	userQuery := &UserList{
		List : ids,
	}

	url := HOST + "/cgi-bin/user/info/batchget?access_token=" + u.token
	b, err := json.Marshal(userQuery)
	if err != nil {
		return nil, err		
	}

	var userList UserListResp
	err = NewRequest().Body(b).JsonResp(&userList)
	if err != nil {
		return nil, err
	}

	return &userList.List, nil
}

