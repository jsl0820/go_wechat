package wechat

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// 设置时间戳
func StampString() string {
	ts64 := time.Now().Unix()
	return strconv.FormatInt(ts64, 10)
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

func NonceStringGenerator(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// //转字符串
// func ToString(value interface{}) string {
//    switch value.(type){
//    case string:
//     	return value.(string)
//    case int32:
// 		v := value.(int)
//     	return strconv.Itoa(v)
//    case int64:
// 		v := value.(int64)
//    		return strconv.FormatInt(v,10)
//    case float32:
// 		v := value.(float32)
//    		return strconv.FormatFloat(v, 'f', -1, 32)
//    case float64:
// 		v := value.(float64)
//    		return strconv.FormatFloat(v, 'f', -1, 64)
//    }
//    return ""
// }

//签名
func StringSign(data map[string]string) string {
	var a []string
	for k, v := range data {
		a = append(a, k+"="+v)
	}
	sort.Strings(a)
	return strings.Join(a, "&")
}

func Md5(singString string) string {
	h := md5.New()
	h.Write([]byte(singString))
	return hex.EncodeToString(h.Sum(nil))
}

//签名
func Sha1Sign(stringSign string) string {
	r := sha1.Sum([]byte(stringSign))
	return hex.EncodeToString(r[:])
}

//map转xml字符串
func MapToXml(data map[string]string) string {
	var xml = `<xml>`
	for k, v := range data {
		xml += `<` + k + `>` + v + `<` + k + `>`
	}
	xml += `</xml>`
	return xml
}

type JsonPare struct {
	Input  string
	Output string
	Data   map[string]string
}

func (js *JsonPare) Write(b []byte) (n int, err error) {
	js.Output += string(b)
	return len(b), nil
}

func (js *JsonPare) Pare() {
	tmpl, err := template.New("json").Parse(js.Input)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(js, js.Data)
	if err != nil {
		panic(err)
	}
}

//
func PareToJson(josnTmpl string, data map[string]string) string {

	js := &JsonPare{
		Data:  data,
		Input: josnTmpl,
	}

	js.Pare()
	return js.Output
}


//构建请求
func Url(url string) string {
	if token, err := TokenInstance.Get(); err != nil {
		panic(err)
		log.Println(err)
	}

	return strings.Replace(url, "{{TOKEN}}", token, -1)
}
