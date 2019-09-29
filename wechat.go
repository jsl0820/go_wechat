package wechat

const HOST = "https://api.weixin.qq.com"
const MCH_HOST = "https://api.mch.weixin.qq.com"

var Wxconfig Config

type Config struct {
	WxAppId     string
	WxAppSecret string
	MpAppid     string
	MpAppsecret string
	PayKey      string
	MchId       string
	Expires     uint
}

func WxConfig(config Config) {
	Wxconfig = config
}

func GetConfig() Config {
	return Wxconfig
}
