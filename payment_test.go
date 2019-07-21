package wechat

import (
	"fmt"
	"testing"
)


func Test_GetPayMent(t *testing.T) {

	option := PayOption {
		AppId : "wxd598f39ca93bd3d3",
		MchId : "1534603911",
		Body : "明科-支付",
		OutTradeNo: "201905141557834400", 
		TotalFee:200,
		NotifyUrl:"http://mingkesy.com/wechat/customized",
		IP : "47.110.72.233", 
		Key: "6IKH14AGWxK92Zqbn1VogXTj8NoAUui2", 
		Openid: "od3875HgOqZ06cx6xqHKadmn9FZc",
	}
	res, err  := SetPayOption(option).GetPayMent()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
	// fmt.Println("result_code", res.ResultCode)
	// fmt.Println("ReturnCode", res.ReturnCode)
}

