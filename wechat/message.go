package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//发送模板信息
const WxApi = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN"
const MpApi = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=ACCESS_TOKEN"


//模板信息配置
type TempMsgConfig struct {
	Touser     string      `json:"touser"`
	TemId string      `json:"template_id"`
	Url        string      `json:"url"`
	Topcolor   string      `json:"topcolor"`
	Mp 
	Data       interface{} `json:"data"`
}

type TempMsgCfg map[string]interface{}


//模板信息返回数据
type TempMsgResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Msgid   string `json:"msgid"`
}

//发送模板信息
func PushTemplateMsg(config TempMsgConfig) TempMsgResponse, error {
	token := TakeToken()
	url := strings.Replace(PushMsgApi, "ACCESS_TOKEN", token.AccessToken, -1)
	data, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}


	reader := bytes.NewReader(data)
	request, err := http.NewRequest("POST", url, reader)


	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		fmt.Println(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(body))

	var msg TempMsgResponse
	json.Unmarshal(body, &msg)
	return msg
}


// 模板消息
type TempMsg struct {
	url string
	plate int  // 0：公众号， 1：小程序
	Config TempMsgCfg
}

//公众号消息
func (t *TempMsg)Wx(cfg TempMsgCfg , token string){
	t.plate = 0 
	t.url = WxApi
} 

//小程序消息
func (t *TempMsg)Mp(cfg TempMsgCfg , token string) TempMsg {
	t.plate = 1
	t.url = MpApi
	return t
}

//发送
func (t *TempMsg)Send(data map[string]string) error {
	t.Config["data"] = data
	cfg, err := json.Marshal(t.Config)

	r := bytes.NewReader(data)
	request, err := http.NewRequest("POST", t.url, r)
}

