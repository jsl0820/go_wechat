package wechat

import (

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

//统一下单接口
type Unified struct {
	PrepayId string
	Resp UnifiedResp
	config map[string]string
	PayInfo map[string]string
}

func(u *Unified)Mp(){
	u.config["appid"] = MpAppid
	u.PayInfo["appid"] = MpAppid
}

func(u *Unified)Wx(){
	u.config["appid"] = WxAppId
	u.PayInfo["appid"] = WxAppId
}

func(u *Unified)PrepayId() string {
	u.config["mch_id"] = MchId
	stringSign = StringSign(p.config) + "&key=" + PayKey
	sign := strings.ToUpper(Md5(stringSign))
	u.config["sign"] :=  sign

	xml := MapToXml(p.config)
	url := "https://api.mch.weixin.qq.com/pay/unifiedorder"

	var resp UnifiedResp
	err := NewRequest(&resp).XmlPost(xml, url)
	u.Resp = resp
	return  resp.PrepayId
}

func(u *Unified)Get()map[string]string{

	u.PayInfo["signType"] = "MD5" 
	u.PayInfo["appId"] = p.config["appid"]
	u.PayInfo["timeStamp"] = StampString()
	u.PayInfo["nonceStr"] = NonceStringGenerator(32)
	u.PayInfo["package"] = "prepay_id=" + res.PrepayId()

	stringSign := StringSign(u.PayInfo) + "&key=" + PayKey
	sign = Md5(stringSign)

	u.PayInfo["paySign"] = strings.ToUpper(sign)

	return u.PayInfo
}

//订单查询
type QueryResp struct {
	ReturnCode string `xml:"ReturnCode"`
	ReturnMsg  string `xml:"return_code"`
	AppId 	   string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"` 
	Sign 	   string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	DeviceInfo string `xml:"device_info"`
	Openid 	   string `xml:"openid"`
	IsSubscribe string `xml:"is_subscribe"`
	TradeType   string `xml:"trade_type"`
	TradeState 	string `xml:"trade_state"`
	BankType 	string `xml:"bank_type"`
	TotalFee 	string `xml:"total_fee"`
	SettlementTotalFee string `xml:"settlement_total_fee"`
	FeeType string `xml:"fee_type"`
	CashFee string `xml:"cash_fee"`
	CashFeeType string `xml:"cash_fee_type"`
	CouponFee string  `xml:"coupon_fee"`
	CouponCount  string `xml:"coupon_count"`
	CouponType string `xml:"coupon_type_$n"`
	CouponId string `xml:"coupon_id_$n"`
	CouponFee string `xml:"coupon_fee_$n"`
	TransactionId string `xml:"transaction_id"`
	UutTradeNo string `xml:"out_trade_no"`
	Attach string `xml:"attach"`
	TimeEnd string `xml:"time_end"`
	TradeStateDesc string `xml:"trade_state_desc"`

}

type Query struct {
	Param map[string]string
}

func (q *Query)config(plat string) error {

	if plat == "mp"{
		q.Param["appid"] = MpAppid
		return nil 
	}

	if plate == "wx" {
		q.Param["appid"]  = WxAppId
		return nil 
	}

	return errors.New("请输入正确的参数！")
}	

func (q *Query)Get(codeType , code string)(QueryResp, error){
	q.Param[codeType] = code
	q.Param["mch_id"] = MchId
	q.Param["nonce_str"] = NonceStringGenerator(32)
	q.Param["sign"] = PaySign(q.Param)

	body := MapToXml(q.Param)
	url := "https://api.mch.weixin.qq.com/pay/orderquery"
	var resp QueryResp
	return NewRequest(&resp).XmlPost(body, url)
}	


//关闭订单
type CloseResp struct {
	ReturnCode string `xml:"ReturnCode"`
	ReturnMsg  string `xml:"return_code"`	
	AppId 	   string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"` 
	Sign 	   string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCodeDes string `xml:"err_code_des"`
}

//退款
type RefundResp struct {
	ReturnCode string `xml:"ReturnCode"`
	ReturnMsg  string `xml:"return_code"`
	ResultCode string `xml:"result_code"`
	ErrCodeDes string `xml:"err_code_des"`
	ErrCode    string `xml:"err_code"`
	AppId 	   string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"` 
	Sign 	   string `xml:"sign"`
	TransactionId string `xml:"transaction_id"`
	OutTradeNo string `xml:"out_trade_no"`
	OutRefundNo string `xml:"out_refund_no"`
	RefundId string `xml:"refund_id"`
	RefundFee	int `xml:"refund_fee"`
	SettlementRefundFee int `xml:"settlement_refund_fee"`
	TotalFee 	string `xml:"total_fee"`
	SettlementTotalFee string `xml:"settlement_total_fee"`
	FeeType string `xml:"fee_type"`
	CashFee string `xml:"cash_fee"`
	CashFeeType string `xml:"cash_fee_type"`
	CashRefundFee string `xml:"cash_refund_fee"`
	CouponType string `xml:"coupon_type_$n"`
	CouponRefundFee string `xml:"coupon_refund_fee"`
	CouponRefundFeeN int `xml:"coupon_refund_fee_$n"`
	CouponCount  string `xml:"coupon_count"`
	CouponRefundId string `xml:"coupon_refund_id_$n"`
}



//支付
type Pay struct {
	
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


//支付签名
func PaySign(param map[string]string) string{
	stringSign = StringSign(p.config) + "&key=" + PayKey
	sign = Md5(stringSign)
	return strings.TouUpper(sign)
}