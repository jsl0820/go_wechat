package mch

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"sort"
	"strings"
	"time"
)


type Param map[string]string

type PayInterface interface {
	nonce(int) string
	sign(map[string]string) string
	md5(string) string
	sha1(string) string
	stringSign(string) string
	init()
}

const CHARS = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Payment struct {
	Resp    interface{}
	Info    map[string]string
	Mchid   string
	WxAppId string
}

//初始化
func (p *Payment) init() {

}

//签名
func (p *Payment) stringSign() string {
	//
	var sings []string
	for k, v := range p.Info {
		sings = append(sings, k+"="+v)
	}

	sort.Strings(sings)
	return strings.Join(sings, "&")
}

//md5加密
func (p *Payment) md5(origin string) string {
	h := md5.New()
	h.Write([]byte(origin))
	return hex.EncodeToString(h.Sum(nil))
}

//sha1加密
func (p *Payment) sha1(origin string) string {
	r := sha1.Sum([]byte(origin))
	return hex.EncodeToString(r[:])
}

//签名
func (p *Payment) sign() string {

}

//时间戳
func (p *Payment) timeStamp() string {

}

//随机字符串
func (p *Payment) nonce(lenght int) string {
	b := []byte(CHARS)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var res []byte
	for i := 0; i < lenght; i++ {
		res = append(res, b[r.Intn(len(b))])
	}

	return string(res)
}
