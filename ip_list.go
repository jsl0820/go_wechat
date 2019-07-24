package wechat

//获取微信的服务器列表

type IpListResp struct {
	ErrCode string   `json:"errcode"`
	ErrMsg  string   `json:"errmsg"`
	IpList  []string `json:"ip_list"`
}

func IpList(token string) (IpListResp, error) {
	url := "https: //api.weixin.qq.com/cgi-bin/getcallbackip?access_token=" + token

	var resp IpListResp
	err := NewRequest(&resp).Get(url)
	if err != nil {
		return nil, err
	}

	return IpListResp, nil
}
