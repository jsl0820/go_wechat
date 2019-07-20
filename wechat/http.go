package wechat

import (
	"encoding/json"
	"bytes"
	"io/ioutil"
	"net/http"
	"errors"

)

func NewRequest(v interface{}) {
	return &Request{
		RespStruct: v
	}
}

type Request struct {
	Response []byte
	RespStruct interface{}
}

func (r *Request) Get(url string) error {
	resp, err := http.Get(url)

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


func (r *Request) Post(url string)  error {
	resp, err := http.Post(url)
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
func (r *Request) XmlPost(xmlStr, url) error {
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
func (r *Request) JsonPost(body interface{}, url) error {
	
	var bodyByte []byte

	switch t := body.(type) {
		case:string
		bodyByte = []byte(body)
		case:[]byte
		bodyByte = body
	default:
		return errors.New("参数类型错误！")
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyByte)))
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