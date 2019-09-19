package wechat

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"

	// "net/url"
	"fmt"
	"log"

	// "path/filepath"
	"mime/multipart"
	"os"
	// "strconv"
)

func NewRequest() *HttpRequest {

	file := make(map[string]string)
	formData := make(map[string]string)

	return &HttpRequest{
		file:     file,
		formData: formData,
	}
}

type HttpRequest struct {
	err         error
	body        io.Reader
	contentType string
	resp        *http.Response
	file        map[string]string
	formData    map[string]string
}

//
func (req *HttpRequest) ContentType(contentType string) *HttpRequest {
	req.contentType = contentType
	return req
}

//断言
func (req *HttpRequest) Body(data interface{}) *HttpRequest {
	switch t := data.(type) {
	case string:
		bf := bytes.NewBufferString(t)
		req.body = ioutil.NopCloser(bf)
	case []byte:
		bf := bytes.NewBuffer(t)
		req.body = ioutil.NopCloser(bf)
	default:
		panic("参数不支持该类型!")
	}

	return req
}

func (req *HttpRequest) Get(uri string) *HttpRequest {
	resp, err := http.Get(uri)
	if err != nil {
		req.err = err
	}

	req.resp = resp
	return req
}

func (req *HttpRequest) Post(uri string) *HttpRequest {
	log.Println("POST", "运行到这里")
	req.contentType = "application/json"
	resp, err := http.Post(uri, "application/json", req.body)

	if err != nil {
		log.Println(err)
		req.resp = resp
	}
	log.Println("POST", "运行到这里2")
	return req
}

//返回Bytes类型
func (req *HttpRequest) Bytes() ([]byte, error) {
	body := req.resp.Body
	respByte, err := ioutil.ReadAll(body)
	if err != nil {
		req.err = err
	}

	defer body.Close()
	return respByte, req.err
}

//返回string类型
func (r *HttpRequest) String() (string, error) {
	b, err := r.Bytes()
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// 保存到
func (r *HttpRequest) SaveTo(path string) error {
	f, err := os.Create(path)
	if err != nil {
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
func (r *HttpRequest) JsonResp(data interface{}) (err error) {
	b, err := r.Bytes()

	if err != nil {
		log.Println(err)
	}

	log.Println("返回数据", string(b))
	return json.Unmarshal(b, data)
}

//返回xml数据解析到结构体
func (r *HttpRequest) XmlResp(data interface{}) (err error) {
	b, err := r.Bytes()
	if err != nil {
		fmt.Println(err)
	}

	return xml.Unmarshal(b, data)
}

//表单上传文件
func (r *HttpRequest) FormFile(field, filename string) *HttpRequest {
	r.file[field] = filename
	return r
}

//表单参数设置
func (r *HttpRequest) FormField(k, v string) *HttpRequest {
	r.formData[k] = v
	return r
}

//上传文件
func (req *HttpRequest) File(field, filename string) *HttpRequest {
	bf := &bytes.Buffer{}
	w := multipart.NewWriter(bf)
	fw, err := w.CreateFormFile(field, filename)
	if err != nil {
		log.Println(err)
	}

	fh, err := os.Open(filename)
	if err != nil {
		req.err = err
	}

	if _, err := io.Copy(fw, fh); err != nil {
		req.err = err
	}

	defer fh.Close()
	req.contentType = w.FormDataContentType()
	w.Close()

	return req
}
