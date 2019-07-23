package wechat

import (
	""
)


type UnifiedResp struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Appid      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	PrepayId   string `xml:"prepay_id"`
	TradeType  string `xml:"trade_type"`
	MwebUrl    string `xml:"mweb_url"`
	CodeUrl    string `xml:"code_url"`
}

//支付
type Pay struct {
	PrepayId string
	Resp UnifiedResp 
	config map[string]string
	PayInfo map[string]string
}

func(p *Pay)Mp(){
	p.config["appid"] = MpAppid
	p.PayInfo["appid"] = MpAppid
}

func(p *Pay)Wx(){
	p.config["appid"] = WxAppId
	p.PayInfo["appid"] = WxAppId
}

func(p *Pay)PrepayId() string {
	p.config["mch_id"] = MchId
	stringSign = StringSign(p.config) + "&key=" + PayKey
	sign := strings.ToUpper(Md5(stringSign))
	p.config["sign"] :=  sign

	xml := MapToXml(p.config)
	url := "https://api.mch.weixin.qq.com/pay/unifiedorder"

	var resp UnifiedResp
	err := NewRequest(&resp).XmlPost(url)
	p.Resp = resp
	return  resp.PrepayId
}


func(p *Pay)Get()map[string]string{

	p.PayInfo["signType"] = "MD5" 
	p.PayInfo["appId"] = p.config["appid"]
	p.PayInfo["timeStamp"] = StampString()
	p.PayInfo["nonceStr"] = NonceStringGenerator(32)
	p.PayInfo["package"] = "prepay_id=" + res.PrepayId()

	stringSign := StringSign(p.PayInfo) + "&key=" + PayKey
	sign = Md5(stringSign)

	p.PayInfo["paySign"] = strings.Upper(sign)

	return p.PayInfo
}

func PayConfig(config map[string]string) Pay {
	return &Pay{sign : config }
}