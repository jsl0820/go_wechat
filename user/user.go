package user

import (
	// "fmt"
	"encoding/json"
	// "strings"
	"errors"
)

type Resp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type Info struct {
	Subscribe      int      `json:"subscribe"`
	OpenId         string   `json:"openid"`
	NickName       string   `json:"nickname"`
	Sex            int      `json:"sex"`
	Language       string   `json:"language"`
	City           string   `json:"city"`
	Province       string   `json:"province"`
	Country        string   `json:"country"`
	HeadImgUrl     string   `json:"headimgurl"`
	SubscribeTime  int64    `json:"subscribe_time"`
	Unionid        string   `json:"unionid"`
	Remark         string   `json:"remark"`
	Groupid        int      `json:"groupid"`
	TagidList      []string `json:"tagid_list"`
	SubscribeScene string   `json:"subscribe_scene"`
	QrScenetr      string   `json:"qr_scene_str"`
}

type User struct {

}


//
func(u *User) Find() *Info {

}


func(u *User) Remark() bool {

}


type BlackList map[string][]string

//黑名单
type BlackListResp struct {
	ErrCode    int       `json:"errcode"`
	ErrMsg     string    `json:"errmsg"`
	Total      int       `json:"total"`
	Count      int       `json:"count"`
	Data       BlackList `json:data`
	NextOpenid string    `json:"next_openid"`
}

type UserList struct {
	List []map[string]string `json:"user_list"`
}

type UserListResp struct {
	List []User `json:"user_info_list"`
}

func NewUsers() *Users {
	t, _ := token.Get()
	return &Users{token: t}
}

//用户
type Users struct {
	token string
}

//用户备注
func (u *Users) Remark(openid, remark string) bool {
	url := HOST + "/cgi-bin/user/info/updateremark?access_token=" + u.token
	body := `{ openid:{{.openid}}, remark:{{.remark}} }`

	var resp UserResp
	err := NewRequest().Body(body).Get(url).JsonResp(&resp)
	if err != nil {
		return false
	}

	if resp.ErrCode != 0 {
		return false
	}

	return true
}

//获取用户信息
func (u *Users) Query(openid string, lang string) (User, error) {
	url := HOST + "/cgi-bin/user/info?access_token="
	url += u.token + "&openid=" + openid + "&lang=" + lang

	var user User
	err := NewRequest().Get(url).JsonResp(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// //批量获取用户信息
func (u *Users) QueryAll(ids []map[string]string) (*[]User, error) {
	userQuery := &UserList{
		List: ids,
	}

	url := HOST + "/cgi-bin/user/info/batchget?access_token=" + u.token
	b, err := json.Marshal(userQuery)
	if err != nil {
		return nil, err
	}

	var userList UserListResp
	err = NewRequest().Body(b).Get(url).JsonResp(&userList)
	if err != nil {
		return nil, err
	}

	return &userList.List, nil
}

// //获取用户地理位置
// func (u *Users)Location(openid, lang string ){
// 	t, _ := token.Get()
// 	url := HOST +"/cgi-bin/user/info?access_token="+ t +"&openid="+ openid +"&lang=" + lang

// }

// //拉黑用户
// func (u *Users)Block(openids ...string)(bool, error){
// 	t, _ := token.Get()
// 	url := HOST + "/cgi-bin/tags/members/batchblacklist?access_token=" + t

// 	if len(openids...) > 20 {
// 		return false, errors.New("数量多，一次最多能操作20个用户")
// 	}

// 	users := strings.Join(openids, ",")
// 	data := `{"openid_list:[" `+ users +`]}`

// 	var resp UserResp
// 	err := NewRequest().Body(data).JsonResp(&resp)
// 	if err != nil {
// 		return false , err
// 	}

// 	if resp.ErrCode == 0 {
// 		return true, nil
// 	}

// 	return false, errors.New("操作失败！Error:" + resp.ErrMsg)
// }

// //拉黑用户
// func (u *Users)CancelBlock(openids ...string)(bool, error){
// 	t, _ := token.Get()
// 	url := HOST + "/cgi-bin/tags/members/batchunblacklist?access_token=" + t
// 	if len(openids...) > 20 {
// 		return false, errors.New("数量多，一次最多能操作20个用户")
// 	}

// 	users := strings.Join(openids, ",")
// 	data := `{"openid_list:[" `+ users +`]}`

// 	var resp UserResp
// 	err := NewRequest().Body(data).JsonResp(&resp)
// 	if err != nil {
// 		return false , err
// 	}

// 	if resp.ErrCode == 0 {
// 		return true, nil
// 	}

// 	return false, errors.New("操作失败！Error:" + resp.ErrMsg)
// }

//获取黑名单
func (u *Users) BlackList(openid string) (BlackListResp, error) {

	var resp BlackListResp
	err := NewRequest().JsonResp(&resp)
	if err != nil {
		return resp, err
	}

	if resp.ErrCode == 0 {
		return resp, nil
	}

	return resp, errors.New("操作失败！Error:" + resp.ErrMsg)
}
