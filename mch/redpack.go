package mch

import (
	. "github.com/jsl0820/wechat"
)

const REPACK_URL = "/mmpaymkttransfers/sendredpack"

type RedPack struct {
	Payment
}

//红包返回数据
type RedPacketResp struct {
	ReturnCode  string `xml:"return_code"`  //必填，返回状态码
	ReturnMsg   string `xml:"return_msg"`   //非必填，返回状态码
	ResultCode  string `xml:"result_code"`  //必填，业务结果
	ErrCode     string `xml:"err_code"`     //非必填，错误代码
	ErrCodeDes  string `xml:"err_code_des"` //非必填，错误代码描述
	MchBillno   string `xml:"mch_billno"`   //必填，商户订单号
	MchId       string `xml:"mch_id"`       //必填，返回商户号
	Wxappid     string `xml:"wxappid"`      //必填，返回公众账号appid
	ReOpenid    string `xml:"re_openid"`    //必填，返回用户openid
	TotalAmount int    `xml:"total_amount"` //必填，付款金额
	SendListid  string `xml:"send_listid"`  //必填，微信单号
}
