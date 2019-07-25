package wechat

import (
	"encoding/json"
	"errors"
)

//发送模板信息
// var tpMsgWxApi = HOST + "/cgi-bin/message/template/send?access_token=ACCESS_TOKEN"
// var tpMsgMpApi = HOST + "/cgi-bin/message/wxopen/template/send?access_token=ACCESS_TOKEN"

type TempMsgCfg map[string]interface{}

//模板信息返回数据
type TempMsgResp struct {
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
	t.url = HOST + "/cgi-bin/message/template/send?access_token=ACCESS_TOKEN"
}

//小程序消息
func (t *TempMsg) Mp(cfg TempMsgCfg, token string)  {
	t.plate = 1
	t.url = HOST + "/cgi-bin/message/wxopen/template/send?access_token=ACCESS_TOKEN"
}

//发送
func (t *TempMsg) Send(data map[string]string) error {
	t.Config["data"] = data
	cfg, err := json.Marshal(t.Config)
	if err != nil {
		return err
	}

	var resp TempMsgResp
	err = NewRequest().Get(t.url).JsonResp(&resp)
	if err != nil {
		return err
	}

	if resp.ErrCode != 0 {
		return errors.New("消息发送失败 msg:" + resp.ErrMsg)
	}

	return nil
}
