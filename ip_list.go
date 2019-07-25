package wechat

//获取微信的服务器列表

type IpListResp struct {
	ErrCode string   `json:"errcode"`
	ErrMsg  string   `json:"errmsg"`
	IpList  []string `json:"ip_list"`
}

func IpList() ([]string, error) {
	t, err  := token.Get()
	url := HOST + "/cgi-bin/getcallbackip?access_token=" + t
	var resp IpListResp
	err = NewRequest().Get(url).JsonResp(&resp)
	
	if err != nil {
		return nil, err
	}

	return resp.IpList, nil
}
