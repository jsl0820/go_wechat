package miniapp

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"strings"
	"net/http"
	"io/ioutil"
	"crypto/aes"
	"crypto/cipher"
)

//code换openid
const getUserIdApi = "https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code"

//发送模板信息
const wxMimiMsgApi = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=ACCESS_TOKEN"

//获取微信access_token
const accessTokenApi = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET"

//小程序的
type AppInfo struct {
	AppId     string
	AppSecret string
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"err_code"`
	ErrMsg      string `json:"err_msg"`
}

//获取openid返回数据
type UserId struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

//模板信息配置
type TempMsgConfig struct {
	FromId          string      `json:"form_id"`
	TemplateId      string      `json:"template_id"`
	Touser          string      `json:"touser"`
	Page            string      `json:"page"`
	Data            interface{} `json:"data"`
	EmphasisKeyword string      `json:"emphasis_keyword"`
}

//模板信息返回数据
type TempMsgResponse struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	TemplateId string `json:"template_id"`
}

type UserInfo struct {
	Openid    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	UnionID   string `json:"unionId"`
	WaterMark `json:"watermark"`
}

type WaterMark struct {
	AppID     string `json:"appid"`
	Timestamp int    `json:"timestamp"`
}

//获取access_token
func GetAccessToken(info AppInfo) *AccessToken {
	appid := info.AppId
	secret := info.AppSecret
	u := strings.Replace(accessTokenApi, "APPID", appid, -1)
	url := strings.Replace(u, "APPSECRET", secret, -1)
	res := httplib.Get(url)
	token := new(AccessToken)
	res.ToJSON(&token)
	return token
}

//发送模板信息
func PushTemplateMsg(info AppInfo, config TempMsgConfig) *TempMsgResponse {
	token := GetAccessToken(info)
	url := strings.Replace(wxMimiMsgApi, "ACCESS_TOKEN", token.AccessToken, -1)
	data, _ := json.Marshal(config)
	res := httplib.Post(url)
	res.Body(string(data))
	msg := new(TempMsgResponse)
	res.ToJSON(&msg)
	return msg
}

//用code换取openid
func GetMpUserId(appid, secret, code string) UserId {
	i := strings.Replace(getUserIdApi, "APPID", appid, -1)
	u := strings.Replace(i, "SECRET", secret, -1)
	url := strings.Replace(u, "JSCODE", code, -1)
	fmt.Println(url)
	// res := httplib.Get(url)
	// user := new(UserId)
	// res.ToJSON(&user)
	// fmt.Println(user)
	// return user
	
	resp, err := http.Get(url)	
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)	
	var user UserId
	fmt.Println(string(body))
	json.Unmarshal(body, &user)
	return user
}


func DecryptData(encryptedData string, iv string , sessionKey string) []byte {

	aseKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		fmt.Println(err)
	}
	aesCipher, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		fmt.Println(err)
	}
	aesIV, err :=  base64.StdEncoding.DecodeString(iv)
	if err != nil {
		fmt.Println(err)
	}
	block, err := aes.NewCipher(aseKey)
	
	if err != nil {
		fmt.Println(err)
	}

	blockMode := cipher.NewCBCDecrypter(block, aesIV)
	decrypted := make([]byte, len(aesCipher))
	blockMode.CryptBlocks(decrypted, aesCipher)
	decrypted = PKCS5UnPadding(decrypted)
	fmt.Println("string", string(decrypted))
	// fmt.Println("userinfo:", info)
	return decrypted
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}