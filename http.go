package wechat

import (
	"encoding/json"
	"bytes"
	"io/ioutil"
	"net/http"
	"errors"

)

func NewRequest(v interface{}) *Request {
	return &HttpRequest{
		RespStruct: v,
	}
}

type HttpRequest struct {
	RespStruct interface{}
	request http.Request
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

func (r *HttpRequest) Get(url string) error {

	resp, err := http.Get(url,)

	if err != nil {
		return  err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return  err
	}
	return json.Unmarshal(b, r.RespStruct)
}


func (r *HttpRequest) Post(url string)  error {

	resp, err := http.Post(url, )
	if err != nil {
		return  err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return  err
	}
	return json.Unmarshal(body,r.RespStruct)
}

//用xml数据请求
func (r *Request) XmlPost(xmlStr, url string) error {
	req, err := http.NewRequest("POST", url, bytes.NewReader([]byte(xmlStr)))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	client := http.Client{}
	resp, err := client.Do(req)
	if (err != nil){
		return err
	}		

	return xml.Unmarshal(respBytes, r.RespStruct)
}


//用xml数据请求
func (r *Request) JsonPost(body interface{}, url string) error {
	
	var bodyByte []byte

	switch body.(type) {
	case string:
		bodyByte = []byte(body)
	case []byte:
		bodyByte = body
	default:
		return errors.New("参数类型错误！")
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyByte))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(req)
	if (err != nil){
		return err
	}		

	return json.Unmarshal(respBytes, r.RespStruct)
}