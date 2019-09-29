package mch

import (
	"errors"
	. "github.com/jsl0820/wechat"
)

const TYPE_1 = "APP -APP"
const TYPE_2 = "JSAPI -JSAPI"
const TYPE_3 = "NATIVE -Native"

const ORDER_QUERY = "/pay/orderquery"
const UNIFIED_ORDER  =  "/pay/unifiedorder"
const CLOSE_ORDER = "/pay/closeorder"
const REFUND = "/secapi/pay/refund"
const REFUND_QUERY = "/pay/refundquery"
const BILL = "/pay/downloadbill"
const FUND_FLOW = "/pay/downloadfundflow"
const QUERY_COMMENT = "/billcommentsp/batchquerycomment"

//统一下单接口返回
type OrderResp struct {
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

type Order struct {
	Payment
}

//生产微信支付订单
func (ord *Order) Create(payType uint8) {
	ord.Info["sign"] = ord.sign()
	ord.Info["mch_id"] = ord.config.MchId
	ord.Info["wxappid"] = ord.config.WxAppId

	url := PAY_HOST + REPACK
	body := ord.xml(ord.Info)

	var resp OrderResp
	request := NewRequest().Body(body)
	request.ContentType("application/xml")
	e := request.Post(url).XmlResp(&resp)
	if e != nil {
		return nil, e
	}

	if resp.ErrCode != 0 {
		return nil, errors.New(resp.ErrCodeDes)
	}

	return resp, nil
}


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

//查询订单
func (ord *Order) Query(*QueryResp, error){
	ord.init()
	url := PAY_HOST + REPACK
	body := ord.xml(ord.Info)

	var resp OrderResp
	request := NewRequest().Body(body)
	request.ContentType("application/xml")
	e := request.Post(url).XmlResp(&resp)
	if e != nil {
		return nil, e
	}

	if resp.ErrCode != 0 {
		return nil, errors.New(resp.ErrCodeDes)
	}

	return resp, nil
}

//关闭订单
type CloseOrderResp struct {
	ReturnCode string `xml:"ReturnCode"`
	ReturnMsg  string `xml:"return_code"`
	AppId      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCodeDes string `xml:"err_code_des"`
}

//关闭订单
func (ord *Order)Close()(*CloseOrderResp, error){
	ord.init()
	var resp CloseOrderResp
	request := NewRequest().Body(body)
	request.ContentType("application/xml")
	e := request.Post(url).XmlResp(&resp)
	if e != nil {
		return nil, e
	}

	if resp.ErrCode != 0 {
		return nil, errors.New(resp.ErrCodeDes)
	}

	return resp, nil
}

