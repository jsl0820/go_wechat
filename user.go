package wechat

type UserResq struct {
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
	UserInfoList []User `json:"user_info_list"` 
}


type Users map[string]interface{}

//获取用户信息
func GetUser(token string, users User) (User, error) {
	url := "https://api.weixin.qq.com/cgi-bin/user/info?access_token="
	url += token + "&openid=" + opneid + "&lang=zh_CN"
	var user User
	err := NewRequest(&user).Get(url)
	if err != nil {
		return nil, err
	}

	return u, nil
}

//批量获取用户信息
func GetUserList(token string, ids []map[string]string) ([]User, error) {
	users := make(Users)
	users["user_list"] = users
	url := "https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=" + token
	b, err := json.Marshal(users)
	if err != nil {
		return nil, err		
	}

	var userList UserList
	err = NewRequest(&userList).JsonPost(url)
	if err != nil {
		return nil, err
	}

	return userList, nil
}

