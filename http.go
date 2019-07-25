package wechat

import (
	"encoding/json"
	"encoding/xml"
	"bytes"
	"io/ioutil"
	"net/http"
	// "errors"
	"net/url"
	"fmt"
	// "log"
)


func NewRequest() *HttpRequest {
	return &HttpRequest{}
}

type HttpRequest struct {
	request http.Request
	response http.Response
}

//请求体
func (r *HttpRequest)Body(data interface{}) *HttpRequest {
	switch t := data.(type){
	case string :
		bf := bytes.NewBufferString(t)
		r.request.Body = ioutil.NopCloser(bf)
		r.request.ContentLength = int64(len(t))
	case []byte:
		bf := bytes.NewBuffer(t)
		r.request.Body = ioutil.NopCloser(bf)
		r.request.ContentLength = int64(len(t))	
	}
	return r
}

//设置header
func (r *HttpRequest)Header(key , value string)*HttpRequest{
	r.request.Header.Set(key, value)
	return r
}

func (r *HttpRequest) Get(urlString string) *HttpRequest {
	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Println(err)
	}

	r.request.URL = u
	r.request.Method = "GET"
	return r
}

func (r *HttpRequest) Post(urlString string)  *HttpRequest {
	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Println(err)
	}

	r.request.URL = u
	r.request.Method = "POST"
	return r
}

//发起请求
func (r *HttpRequest) do() (*http.Response, error) {
	clinet := &http.Client{}
	return clinet.Do(&r.request)
}

//返回Bytes类型
func (r *HttpRequest)Bytes()([]byte, error){
	resp, err := r.do()
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//
func(r *HttpRequest)JsonResp(data interface{}) (err error) {
	b, err := r.Bytes()
	if err != nil {
		fmt.Println(err)
	}

	return json.Unmarshal(b, data)
}

//
func(r *HttpRequest)XmlResp(data interface{}) (err error) {
	b, err := r.Bytes()
	if err != nil {
		fmt.Println(err)
	}

	return xml.Unmarshal(b, data)
}


