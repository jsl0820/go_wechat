package mch

import "testing"

func TestCreate(t *testing.T) {
	info := Param{
		"body" : "",
		"detail" : "",  //非必须
		"attach" : "",  //非必须
		"out_trade_no":"",
		"fee_type":"", //非必须
		"total_fee" :"" ,
		"spbill_create_ip":"",
		"time_start" :  "", //非必须
		"time_expire" : "", //非必须
		"goods_tag" : "" , //非必须
		"notify_url" :"",
		"trade_type" : "",
		"product_id" : "", //trade_type=NATIVE时，此参数必传。
		"limit_pay" : "",  // 上传此参数no_credit--可限制用户不能使用信用卡支付
		"openid" : "",  // trade_type=JSAPI时（即JSAPI支付），此参数必传，
		"receipt" : "",
		"scene_info" : "",
	}

	ord := &Order{Info:info}
	ord.Create()
}

func TestQuery(t *testing.T) {
	info :=
}


