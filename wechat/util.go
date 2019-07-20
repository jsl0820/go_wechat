package wechat

import (
	"time"
	"strconv"
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

//转字符串
func ToString(value interface{}) string {
   switch t := value.(type){
   case: string
    	return value   
   case: int32
    	return strconv.Itoa(value)	
   case: int64
   		return strconv.FormatInt(int64,10)
   case: float32
   		return strconv.FormatFloat(value, 'f', -1, 32)
   case: float64	
   		return strconv.FormatFloat(value, 'f', -1, 64)		 		   
   }
}
