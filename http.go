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
	"log"
	// "path/filepath"
	"mime/multipart"
	"os"
	"io"
	// "strconv"
)


func NewRequest() *HttpRequest {

	files := make(map[string]string)
	formData := make(map[string]string)
	request := http.Request {
		Header:make(http.Header),
	}

	return &HttpRequest{
		files : files,
		formData : formData,
		request:request,
	}
}

type HttpRequest struct {
	files map[string]string
	formData map[string]string
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
		log.Println(err)
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
	log.Println("请求头", r.request.Header)
	log.Println("请求链接", r.request.URL)

	return clinet.Do(&r.request)
}

//返回Bytes类型
func (r *HttpRequest)Bytes()([]byte, error){
	resp, err := r.do()

	log.Println(resp.StatusCode)
	// log.Println(resp.Body)
	log.Println(err)

	if err != nil {
		log.Println("请求出错!")
		panic(err)
	}

	// log.Println("请求", r.request)
	// log.Println("响应", resp.Header)

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//返回string类型
func (r *HttpRequest)String()(string, error){
	b, err := r.Bytes()
	if err != nil {
		return "" , err
	}

	return string(b), nil
}

// 保存到
func (r *HttpRequest)SaveTo(path string)error{
	f, err := os.Create(path)
	if err != nil{
		return err
	} 

	defer f.Close()	
	b, err := r.Bytes()
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	return err 
}

//返回json数据解析到结构体
func(r *HttpRequest)JsonResp(data interface{}) (err error) {
	b, err := r.Bytes()

	if err != nil {
		log.Println(err)
	}

	return json.Unmarshal(b, data)
}

//返回xml数据解析到结构体
func(r *HttpRequest)XmlResp(data interface{}) (err error) {
	b, err := r.Bytes()
	if err != nil {
		fmt.Println(err)
	}

	return xml.Unmarshal(b, data)
}


//表单上传文件
func(r *HttpRequest)FormFile(field, filename string) *HttpRequest{
	r.files[field] = filename
	return r
}

//表单参数设置
func(r *HttpRequest)Param(k, v string) *HttpRequest{
	r.formData[k] = v
	return r
}

//构建表单
//这里是有文件的POST表单
func(r *HttpRequest)Form() *HttpRequest{
	//读取文件
	if len(r.files) > 0 {
		bf := &bytes.Buffer{}
		w := multipart.NewWriter(bf)
		defer w.Close()
		for field, fileName := range r.files {
			fw, err := w.CreateFormFile(field, fileName)
			// fw, err := w.CreateFormFile(field, filepath.Base(fileName))
			if err != nil {
				log.Println(err)
			}
			
			fh, err := os.Open(fileName)
			if err != nil {
				panic(err)
			}
			
			defer fh.Close()
			io.Copy(fw, fh)
		}
		// length :=  strconv.Itoa(bf.Len())
		r.Header("Content-Type", w.FormDataContentType())
		// r.Header("Content-Length", length)
		for k, v := range r.formData {
			w.WriteField(k, v)
		}

		r.request.Body =  ioutil.NopCloser(bf)
	} 
	return r
}


//构建表单
//这里是有文件的POST表单
func(r *HttpRequest)PostFile(url, field, filename string) {

	bf := &bytes.Buffer{}
	w := multipart.NewWriter(bf)
	defer w.Close()

	log.Println(field)

	fw, err := w.CreateFormFile(field, filename)
	if err != nil {
		log.Println(err)
	}
	
	fh, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	
	defer fh.Close()

	io.Copy(fw, fh)
	contentType :=  w.FormDataContentType()

	// r.Header("Content-Length", length)
	// for k, v := range r.formData {
	// 	w.WriteField(k, v)
	// }

	resp, err := http.Post(url, contentType, bf)
	// r.request.Body =  ioutil.NopCloser(bf)
	
	// log.Println("响应1",resp)	

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println(err)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("响应", string(respBody))
}


