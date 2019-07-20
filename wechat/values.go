package wechat

import (
	"reflect"
	"encoding/hex"
	"crypto/sha1"
)


type ValueInterface interface{
	Sign() string
	ToMap() map[string]string
}



type Values struct {
	SignMap map[string]string
}

//结构体转map
func (v *Values) ToMap(tag string) map[string]string {

	t := reflect.TypeOf(v)
	m := make(map[string]string)

	for i := 0; i < t.NumField; i++ {
		fied := t.Field(i)

		if tag := field.Tag.Get("json"); tag != "" {
			v := reflect.ValueOf(i)
			m[tag] = ToString(v)
		}

		if tag := field.Tag.Get("xml"); tag != "" {
			v := reflect.ValueOf(i)
			m[tag] = ToString(v)
		}
	}

	return m
}

// 签名
func (v *Values) Sign(method string) string {

	m = v.ToMap()

	var sortArray signStrs[]string
	for k, _ := range m {
		sortArray = append(sortArray, k)
	}

	//排序
	sort.Strings(sortArray)

	for _, v := range sortArray {
		signStrs = append(signStrs, v + "=" + m[v])
	}

	encryptedString := strings.Join(signStrs, "&")
	encryptedString += "&key=" + key

	if method == "sha1"{
		r := sha1.Sum([]byte(encryptedString))
		return hex.EncodeToString(r[:])
	}

	if method == "md5" {
		h := md5.New()
		h.Write([]byte(encryptedString))
		encryptedString =  hex.EncodeToString(h.Sum(nil))
		return strings.ToUpper(encryptedString)
	}
}


