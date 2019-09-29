package mch

import (
	"errors"
	. "github.com/jsl0820/wechat"
)

const MP_REDPACK = 1
const WX_REDPACK = 2
const MP_REPACK = "/mmpaymkttransfers/sendminiprogramhb"

const REPACK = "/mmpaymkttransfers/sendredpack"
const GROUP_REDBACK = "/mmpaymkttransfers/sendgroupredpack"
const REDBACK_RECD = "/mmpaymkttransfers/gethbinfo"

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

type RedPack struct {
	Payment
}

//发送红包
func (red *RedPack) SendRedPack() (*RedPacketResp, error) {
	red.Info["sign"] = red.sign()
	red.Info["mch_id"] = red.config.MchId
	red.Info["wxappid"] = red.config.WxAppId
	red.Info["nonce_str"] = red.nonce()

	url := PAY_HOST + REPACK
	body := red.xml(red.Info)

	var resp RedPacketResp
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

//裂变红包
func (red *RedPack) SendGoupredPack() {
	red.Info["sign"] = red.sign()
	red.Info["mch_id"] = red.config.MchId
	red.Info["wxappid"] = red.config.WxAppId
	red.Info["nonce_str"] = red.nonce()

	url := PAY_HOST + REPACK
	body := red.xml(red.Info)

	var resp RedPacketResp
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

//红包记录
func (red *RedPack) Record(cat uint) {
	red.Info["sign"] = red.sign()
	red.Info["mch_id"] = red.config.MchId
	red.Info["wxappid"] = red.config.WxAppId
	red.Info["nonce_str"] = red.nonce()

	url := PAY_HOST + REPACK
	if cat == MP_REDPACK {
		url = PAY_HOST + MP_REPACK
	}

	body := red.xml(red.Info)

	var resp RedPacketResp
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
