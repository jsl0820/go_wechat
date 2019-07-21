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

type TempMsgCfg map[string]interface{}

//模板信息返回数据
type TempMsgResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Msgid   string `json:"msgid"`
}

// 模板消息
type TempMsg struct {
	url    string
	plate  int // 0：公众号， 1：小程序
	Config TempMsgCfg
}

//公众号消息
func (t *TempMsg) Wx(cfg TempMsgCfg, token string) {
	t.plate = 0
	t.url = WxApi
}

//小程序消息
func (t *TempMsg) Mp(cfg TempMsgCfg, token string) TempMsg {
	t.plate = 1
	t.url = MpApi
	return t
}

//发送
func (t *TempMsg) Send(data map[string]string) error {
	t.Config["data"] = data
	cfg, err := json.Marshal(t.Config)
	if err != nil {
		return err
	}

	var resp TempMsgResponse
	err := NewRequest(&resp).JsonPost(t.url)
	if err != nil {
		return err
	}

	if resp.ErrCode != 0 {
		return errors.New("消息发送失败 msg:" + resp.ErrMsg)
	}

	return nil
}
