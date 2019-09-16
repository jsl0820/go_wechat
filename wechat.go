package wechat

// var WxAppId string
// var WxAppSecret string
// var MpAppid string
// var MpAppsecret string
// var MchId string
// var PayKey string
const HOST = "https://api.weixin.qq.com"

var Wxconfig Config

type Config struct {
	WxAppId     string
	WxAppSecret string
	MpAppid     string
	MpAppsecret string
	PayKey      string
	MchId       string
}

// func init() {
// 	// WxAppId = "wx582ef3694f7a7546"
// 	// WxAppSecret = "148ee9063222674ef03e4c21776e02cd"
// 	// MpAppid = "wxd598f39ca93bd3d3"
// 	// MpAppsecret = "4cc8ce9de598c2fcfd3dce9d809cb585"
// 	// MchId = ""  //商户号
// 	// PayKey = "" //支付秘钥
// }

func WxConfig(config Config) {
	Wxconfig = config
}
