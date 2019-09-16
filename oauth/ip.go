package oauth

import (
	. "github.com/jsl0820/wechat"
)

const CALLBACK_IPS = "/cgi-bin/getcallbackip?access_token={{TOKEN}}"

//获取微信的服务器列表
type IpListResp struct {
	ErrCode string   `json:"errcode"`
	ErrMsg  string   `json:"errmsg"`
	IpList  []string `json:"ip_list"`
}

func IpList() ([]string, error) {
	url := Url(CALLBACK_IPS)
	var resp IpListResp
	err = NewRequest().Get(url).JsonResp(&resp)
	if err != nil {
		return nil, err
	}

	return resp.IpList, nil
}
