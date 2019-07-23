package wechat

import (
	"time"
	"sort"
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/xml"
	"strconv"
	"crypto/md5"
	"encoding/hex"
	// "net/url"
	"strings"
	"math/rand"
	"errors"
	"crypto/sha1"
)

// feture
// 这个库还需要一个可以处理xml/map/struct之间的转换关系的辅助库

const UnifiedOrderUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"

// 统一下单
type UnifiedOrder struct {
	Appid            string `xml:"appid"`            //公众账号ID
	Body             string `xml:"body"`             //商品描述
	MchId            string `xml:"mch_id"`           //商户号
	NonceStr         string `xml:"nonce_str"`        //随机字符串
	NotifyUrl        string `xml:"notify_url"`       //通知地址
	TradeType        string `xml:"trade_type"`       //交易类型
	SpbillCreateIp   string `xml:"spbill_create_ip"` //支付提交用户端ip
	TotalFee         int    `xml:"total_fee"`        //总金额
	OutTradeNo     	 string `xml:"out_trade_no"`     //商户订单号
	Sign             string `xml:"sign"`             //签名
	Openid  		 string `xml:"openid"`	
}

type UnifiedOrderResponse struct {
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

type Payment struct{
	TimeStamp  	string `json:"timeStamp"`
	NonceStr	string `json:"nonceStr"`
	Package     string `json:"package"`
	SignType    string `json:"signType"`
	PaySign		string `json:"paySign"`
} 


type NotifyResp struct {
	ReturnCode    string  `xml:"return_code"`
	ReturnMsg     string  `xml:"return_msg"`
	Appid         string  `xml:"appid"`
	Mchid         string  `xml:"mch_id"`
	NonceStr      string  `xml:"nonce_str"`
	Sign          string  `xml:"sign"`
	ResultCode    string  `xml:"result_code"`
	Openid        string  `xml:"openid"`
	IsSubscribe   string  `xml:"is_subscribe"`
	TradeType     string  `xml:"trade_type"`
	BankType      string  `xml:"bank_type"`
	TotalFee      string `xml:"total_fee"`
	FeeType       string  `xml:"fee_type"`
	CashFee       int     `xml:"cash_fee"`
	CashFeeType   string  `xml:"cash_fee_type"`
	TransactionId string  `xml:"transaction_id"`
	OutTradeNo    string  `xml:"out_trade_no"`
	TimeEnd       string  `xml:"time_end"`
}



//支付参数
type  PayOption struct {
	AppId string
	MchId string
	Body string
	OutTradeNo string
	TotalFee int
	IP string
	NotifyUrl string
	Key string
	Openid string
}

type SignMap map[string]string

type MiniAppPay struct {
	Option  PayOption
	Unified UnifiedOrder
}

// 设置
func SetPayOption(option PayOption) *MiniAppPay {

	sign := make(SignMap)

	nonceStr := CreateNonceStr(32)

	sign["appid"] =  option.AppId
	sign["mch_id"] = option.MchId
	sign["nonce_str"] = nonceStr
	sign["body"] = option.Body
	sign["out_trade_no"] = option.OutTradeNo
	sign["total_fee"] = strconv.Itoa(option.TotalFee)
	sign["spbill_create_ip"] = option.IP
	sign["notify_url"] = option.NotifyUrl
	sign["trade_type"] = "JSAPI"
	sign["openid"] = option.Openid

	unifiedInfo := UnifiedOrder {
		Appid : option.AppId,
		MchId : option.MchId, 
		NonceStr : nonceStr, 
		Body :  option.Body, 
		OutTradeNo : option.OutTradeNo,
		TotalFee : option.TotalFee,
		SpbillCreateIp : option.IP,
		NotifyUrl : option.NotifyUrl,
		TradeType : "JSAPI",
		Sign : CreateSign(sign, option.Key),
		Openid : option.Openid,
	}

	p := new (MiniAppPay)
	p.Option = option
	p.Unified = unifiedInfo

	fmt.Println("nonceStr:", nonceStr)
	fmt.Println("sign:", sign)
	fmt.Println("unified签名", CreateSign(sign, option.Key))
	return p
}

// 发送请求
func (p *MiniAppPay) GetUnifieldInfo() (UnifiedOrderResponse, error) {
	var err error
	xm, err := xml.Marshal(p.Unified)
	xmlStr := strings.Replace(string(xm), "UnifiedOrder", "xml", -1)
	fmt.Println(xmlStr)
	request, err := http.NewRequest("POST", UnifiedOrderUrl , bytes.NewReader([]byte(xmlStr)))

	request.Header.Set("Accept", "application/xml")
	request.Header.Set("Content-Type", "application/xml;charset=utf-8")

	client := http.Client{}
	response, _ := client.Do(request)
	respBytes, err := ioutil.ReadAll(response.Body)

	var ret UnifiedOrderResponse
	xml.Unmarshal(respBytes, &ret)

	if ret.ResultCode == "FAIL" || ret.ReturnCode == "FAIL"{
		info := "ResultCode:" + ret.ResultCode + ", ReturnCode:"+ ret.ReturnCode
		err = errors.New(info)
		fmt.Println("统一下单接口：", ret)
		fmt.Println("ResultCode", ret.ResultCode)
		fmt.Println("ReturnCode", ret.ReturnCode)
	} 

	return ret, err
}

// 获取支付信息
func (p *MiniAppPay)GetPayMent() (Payment, error) {
	res, err := p.GetUnifieldInfo()
	// fmt.Println(res)	
	sign := make(SignMap)
	sign["appId"] = p.Option.AppId
	sign["timeStamp"] = GetStampString()
	sign["nonceStr"] = CreateNonceStr(32)
	sign["package"] = "prepay_id="+res.PrepayId
	sign["signType"] = "MD5"


	payInfo := Payment{
		TimeStamp : GetStampString(),
		NonceStr : 	CreateNonceStr(32),
		Package : "prepay_id="+res.PrepayId,
		SignType : "MD5",
		PaySign : CreateSign(sign, p.Option.Key),
	}

	fmt.Println(res.PrepayId)

	return payInfo, err 		
}

// 设置签名
func CreateSign(sign SignMap, key string) string {
	// 1.字典序排列, 拼接字符串
	var arr []string
	for k, _ := range sign {
		arr = append(arr, k)
	}
	sort.Strings(arr)
	var sa []string

	fmt.Println(arr)

	for _, v := range arr{
		if sign[v] != ""{
			// str := v + "=" + url.QueryEscape(sign[v])
			str := v + "=" + sign[v]
			sa = append(sa, str)
		}
	}
	strs := strings.Join(sa, "&")
	//2.拼接密钥
	strs = strs + "&key=" + key

	fmt.Println("签名string:", strs)

	h := md5.New()
	h.Write([]byte(strs))

	strs =  hex.EncodeToString(h.Sum(nil))

	fmt.Println("strs:", strs)
	return strings.ToUpper(strs)
}

func Sha1Sign(sign SignMap) string {
	var arr []string
	for k, _ := range sign {
		arr = append(arr, k)
	}
	sort.Strings(arr)
	var sa []string
	for _, v := range arr{
		if sign[v] != ""{
			// str := v + "=" + url.QueryEscape(sign[v])
			str := v + "=" + sign[v]
			sa = append(sa, str)
		}
	}
	strs := strings.Join(sa, "&")
	fmt.Println("签名string:", strs)
	r := sha1.Sum([]byte(strs))
	return hex.EncodeToString(r[:])
}


// 设置时间戳
func GetStampString() string{
	ts64 := time.Now().Unix()
	return strconv.FormatInt(ts64,10)
}

// 随机字符串
func CreateNonceStr(length int) string {
   str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
   bytes := []byte(str)
   result := []byte{}
   r := rand.New(rand.NewSource(time.Now().UnixNano()))
   for i := 0; i < length; i++ {
      result = append(result, bytes[r.Intn(len(bytes))])
   }
   return string(result)
}

// future 
// 把响应转为目标结构体
func XmlRespTo() {
	
}