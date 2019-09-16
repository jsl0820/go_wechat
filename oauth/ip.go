package oauth

import (
	. "github.com/jsl0820/wechat"
)

const CALLBACK_IPS = "/cgi-bin/getcallbackip?access_token={{TOKEN}}"

func IP() *IPResp {
	return &IPResp{}
}

//获取微信的服务器列表
type IPResp struct {
	ErrCode string   `json:"errcode"`
	ErrMsg  string   `json:"errmsg"`
	IpList  []string `json:"ip_list"`
}

func (ips *IPResp) List() ([]string, error) {
	url := Url(CALLBACK_IPS)
	err := NewRequest().Get(url).JsonResp(ips)
	if err != nil {
		return nil, err
	}

	return ips.IpList, nil
}

//IP地址是否在服务器列表中
func (ips *IPResp) Has(ip string) bool {
	ips.List()
	for i := 0; i < len(ips.IpList); i++ {
		if ips.IpList[i] == ip {
			return true
		}
	}

	return false
}
