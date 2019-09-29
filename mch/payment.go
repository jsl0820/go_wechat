package mch

import (
	"errors"
	"fmt"
	"strings"
)

const PAY_HOST = "https://api.mch.weixin.qq.com"

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
	Resp    UnifiedResp
	Param   map[string]string
	PayInfo map[string]string
}

func (u *Unified) Mp() {
	u.Param["appid"] = Wxconfig.MpAppid
	u.PayInfo["appid"] = Wxconfig.MpAppid
}

func (u *Unified) Wx() {
	u.Param["appid"] = Wxconfig.WxAppId
	u.PayInfo["appid"] = Wxconfig.WxAppId
}

// func(u *Unified)PrepayId() string {
// 	u.Param["mch_id"] = MchId
// 	stringSign := StringSign(u.Param) + "&key=" + PayKey
// 	sign := strings.ToUpper(Md5(stringSign))
// 	u.Param["sign"] =  sign

// 	xml := MapToXml(u.Param)
// 	url := HOST + "/pay/unifiedorder"

// 	var resp UnifiedResp
// 	err := NewRequest().Get(url).XmlResp(&resp)
// 	u.Resp = resp
// 	return  resp.PrepayId
// }

// func(u *Unified)Get()map[string]string{

// 	u.PayInfo["signType"] = "MD5"
// 	u.PayInfo["appId"] = u.Param["appid"]
// 	u.PayInfo["timeStamp"] = StampString()
// 	u.PayInfo["nonceStr"] = NonceStringGenerator(32)
// 	u.PayInfo["package"] = "prepay_id=" + u.PrepayId()

// 	stringSign := StringSign(u.PayInfo) + "&key=" + PayKey
// 	sign := Md5(stringSign)

// 	u.PayInfo["paySign"] = strings.ToUpper(sign)

// 	return u.PayInfo
// }

//订单查询
type QueryResp struct {
	ReturnCode         string `xml:"ReturnCode"`
	ReturnMsg          string `xml:"return_code"`
	AppId              string `xml:"appid"`
	MchId              string `xml:"mch_id"`
	NonceStr           string `xml:"nonce_str"`
	Sign               string `xml:"sign"`
	ResultCode         string `xml:"result_code"`
	ErrCode            string `xml:"err_code"`
	ErrCodeDes         string `xml:"err_code_des"`
	DeviceInfo         string `xml:"device_info"`
	Openid             string `xml:"openid"`
	IsSubscribe        string `xml:"is_subscribe"`
	TradeType          string `xml:"trade_type"`
	TradeState         string `xml:"trade_state"`
	BankType           string `xml:"bank_type"`
	TotalFee           string `xml:"total_fee"`
	SettlementTotalFee string `xml:"settlement_total_fee"`
	FeeType            string `xml:"fee_type"`
	CashFee            string `xml:"cash_fee"`
	CashFeeType        string `xml:"cash_fee_type"`
	CouponFee          string `xml:"coupon_fee"`
	CouponCount        string `xml:"coupon_count"`
	CouponType         string `xml:"coupon_type_$n"`
	CouponId           string `xml:"coupon_id_$n"`
	TransactionId      string `xml:"transaction_id"`
	UutTradeNo         string `xml:"out_trade_no"`
	Attach             string `xml:"attach"`
	TimeEnd            string `xml:"time_end"`
	TradeStateDesc     string `xml:"trade_state_desc"`
}

type Query struct {
	appid string
	Param map[string]string
}

func (q *Query) config(plat string) error {

	if plat == "mp" {
		q.Param["appid"] = Wxconfig.MpAppid
		return nil
	}

	if plat == "wx" {
		q.Param["appid"] = Wxconfig.WxAppId
		return nil
	}

	return errors.New("请输入正确的参数！")
}

func (q *Query) Get(codeType, code string) (*QueryResp, error) {
	q.Param["appid"] = q.appid
	q.Param[codeType] = code
	q.Param["mch_id"] = Wxconfig.MchId
	q.Param["nonce_str"] = NonceStringGenerator(32)
	q.Param["sign"] = PaySign(q.Param)

	body := MapToXml(q.Param)
	url := PAY_HOST + "/pay/orderquery"
	var resp QueryResp
	err := PayRequest(url, body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

//关闭订单
type CloseResp struct {
	ReturnCode string `xml:"ReturnCode"`
	ReturnMsg  string `xml:"return_code"`
	AppId      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCodeDes string `xml:"err_code_des"`
}

//退款
type RefundResp struct {
	ReturnCode          string `xml:"ReturnCode"`
	ReturnMsg           string `xml:"return_code"`
	ResultCode          string `xml:"result_code"`
	ErrCodeDes          string `xml:"err_code_des"`
	ErrCode             string `xml:"err_code"`
	AppId               string `xml:"appid"`
	MchId               string `xml:"mch_id"`
	NonceStr            string `xml:"nonce_str"`
	Sign                string `xml:"sign"`
	TransactionId       string `xml:"transaction_id"`
	OutTradeNo          string `xml:"out_trade_no"`
	OutRefundNo         string `xml:"out_refund_no"`
	RefundId            string `xml:"refund_id"`
	RefundFee           int    `xml:"refund_fee"`
	SettlementRefundFee int    `xml:"settlement_refund_fee"`
	TotalFee            string `xml:"total_fee"`
	SettlementTotalFee  string `xml:"settlement_total_fee"`
	FeeType             string `xml:"fee_type"`
	CashFee             string `xml:"cash_fee"`
	CashFeeType         string `xml:"cash_fee_type"`
	CashRefundFee       string `xml:"cash_refund_fee"`
	CouponType          string `xml:"coupon_type_$n"`
	CouponRefundFee     string `xml:"coupon_refund_fee"`
	CouponRefundFeeN    int    `xml:"coupon_refund_fee_$n"`
	CouponCount         string `xml:"coupon_count"`
	CouponRefundId      string `xml:"coupon_refund_id_$n"`
}

type RefundQueryResp struct {
	ReturnCode          string `xml:"ReturnCode"`
	ReturnMsg           string `xml:"return_code"`
	ErrCode             string `xml:"err_code"`
	ErrCodeDes          string `xml:"err_code_des"`
	AppId               string `xml:"appid"`
	MchId               string `xml:"mch_id"`
	NonceStr            string `xml:"nonce_str"`
	Sign                string `xml:"sign"`
	TotalRefundCount    int    `xml:"total_refund_count"`
	TransactionId       string `xml:"transaction_id"`
	OutTradeNo          string `xml:"out_trade_no"`
	TotalFee            string `xml:"total_fee"`
	SettlementTotalFee  string `xml:"settlement_total_fee"`
	FeeType             string `xml:"fee_type"`
	CashFee             string `xml:"cash_fee"`
	RefundCount         int    `xml:"refund_count"`
	OutRefundNo         string `xml:"out_refund_no"`
	RefundId            string `xml:"refund_id"`
	RefundChannelNn     string `xml:"refund_channel_$n"`
	RefundFee           int    `xml:"refund_fee"`
	SettlementRefundFee int    `xml:"settlement_refund_fee"`
	CouponTypeNM        string `xml:"coupon_type_$n_$m"`
	CouponRefundFeeN    int    `xml:"coupon_refund_fee_$n"`
	CouponRefundCountN  int    `xml:"coupon_refund_count_$n"`
	CouponRefundIdNM    string `xml:"coupon_refund_id_$n_$m"`
	CouponRefundFeeNM   string `xlm:"coupon_refund_fee_$n_$m"`
	RefundStatusN       string `xml:"refund_status_$n"`
	RefundAccountN      string `xml:"refund_account_$n"`
	RefundRecvAccoutN   string `xml:"refund_recv_accout_$n"`
	RefundSuccessTimeN  string `xml:"refund_success_time_$n"`
}

//退款
type Refund struct {
	appid     string
	mchId     string
	orderType string
	orderNum  string
}

//申请退款
func (r *Refund) Apply(info map[string]string) RefundQueryResp {
	info[r.orderType] = r.orderNum
	info["appid"] = r.appid
	info["mch_id"] = r.mchId
	info["nonce_str"] = NonceStringGenerator(32)
	info["sign"] = PaySign(info)

	body := MapToXml(info)
	url := "https://api.mch.weixin.qq.com/secapi/pay/refund"

	var resp RefundQueryResp
	err := PayRequest(url, body, &resp)
	if err != nil {
		fmt.Println(err)
	}

	return resp
}

//查询退款
func (r *Refund) Query(info map[string]string) RefundQueryResp {
	var resp RefundQueryResp
	body := MapToXml(info)
	url := "https://api.mch.weixin.qq.com/pay/refundquery"

	err := PayRequest(url, body, &resp)
	if err != nil {
		fmt.Println(err)
	}

	return resp
}

//下载对账单
type DownLoadBillResp struct {
	ReturnCode string `xml:"ReturnCode"`
	ReturnMsg  string `xml:"return_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
}

//账单
type Bill struct {
	Param map[string]string
	Resp  DownLoadBillResp
}

//
func (b *Bill) Get(billType, date string) *Bill {
	b.Param["bill_date"] = date
	b.Param["bill_type"] = billType
	b.Param["sign"] = PaySign(b.Param)

	url := "https://api.mch.weixin.qq.com/pay/downloadbill"
	body := MapToXml(b.Param)

	req := NewRequest().Body(body).Post(url)
	req.ContentType("application/xml;charset=utf-8")

	var resp DownLoadBillResp
	err := req.XmlResp(&resp)
	if err != nil {
		fmt.Println(err)
	}

	return b
}

//保存
func (b *Bill) SaveAs(path string, fileName string) {

}

//支付
type Pay struct {
	appid string
}

//统一下单
func (p *Pay) Unified(param map[string]string) *Unified {
	return &Unified{Param: param}
}

//订单查询
func (p *Pay) Query(codeType, code string) *Query {
	param := make(map[string]string)
	param["mch_id"] = Wxconfig.MchId
	param["nonce_str"] = NonceStringGenerator(32)
	param[codeType] = code
	return &Query{Param: param}
}

//关闭订单
func (p *Pay) Close(codeType, code string) CloseResp {
	param := make(map[string]string)
	param["mch_id"] = Wxconfig.MchId
	param["nonce_str"] = NonceStringGenerator(32)
	param[codeType] = code
	param["sign"] = PaySign(param)
	body := MapToXml(param)
	url := "https://api.mch.weixin.qq.com/pay/closeorder"

	var resp CloseResp
	req := NewRequest().Body(body).Post(url)
	req.ContentType("application/xml;charset=utf-8")
	err := req.XmlResp(&resp)
	if err != nil {
		fmt.Println(err)
	}

	return resp
}

//申请退款
func (p *Pay) Refund(codeType, code string) RefundResp {
	param := make(map[string]string)

	param["appid"] = p.appid
	param["mch_id"] = Wxconfig.MchId
	param["nonce_str"] = NonceStringGenerator(32)
	param[codeType] = code

	body := MapToXml(param)
	url := "https://api.mch.weixin.qq.com/pay/closeorder"
	var resp RefundResp

	err := PayRequest(url, body, &resp)
	if err != nil {
		fmt.Println(err)
	}

	return resp
}

//下载对账单
func (p *Pay) Bill() *Bill {
	param := make(map[string]string)
	param["appid"] = p.appid
	param["mch_id"] = Wxconfig.MchId
	param["nonce_str"] = NonceStringGenerator(32)

	return &Bill{Param: param}
}

//支付签名
func PaySign(param map[string]string) string {
	stringSign := StringSign(param) + "&key=" + Wxconfig.PayKey
	sign := Md5(stringSign)
	return strings.ToUpper(sign)
}

func NewPay(plat string) *Pay {

	var appid string
	if plat == "wx" {
		appid = Wxconfig.WxAppId
	}

	if plat == "mp" {
		appid = Wxconfig.MpAppid
	}

	return &Pay{appid: appid}
}

//请求
func PayRequest(url string, body interface{}, resp interface{}) error {
	req := NewRequest().Body(body).Post(url)
	req.ContentType("application/xml;charset=utf-8")
	err := req.XmlResp(resp)
	if err != nil {
		return err
	}

	return nil
}
