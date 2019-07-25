package wechat

var WxAppId string
var WxAppSecret string
var MpAppid string
var MpAppsecret string
var MchId string
var PayKey string
const  HOST = "https://api.weixin.qq.com"

func init() {
	WxAppId = "wx75d0a800a00671a1"
	WxAppSecret = "de3426ea07a05887a220c91232fcc9e7"
	MpAppid = "wxd598f39ca93bd3d3"
	MpAppsecret = "4cc8ce9de598c2fcfd3dce9d809cb585"
	MchId = ""  //商户号
	PayKey = "" //支付秘钥
}
