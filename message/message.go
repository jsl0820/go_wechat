package message

import (
	"encoding/json"
	"errors"
	"fmt"
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

type TempInfo struct {
	TemplateId      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
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

//设置所在行业
func (t *TempMsg) SetIndustry(idst1, idst2 string) (bool, error) {
	tk, err := token.Get()
	if err != nil {
		return false, err
	}

	body := `{ industry_id1:{{.idst1}}, industry_id2:{{.idst2}} }`
	url := HOST + "/cgi-bin/template/api_set_industry?access_token=" + tk

	var resp TempMsgResp
	err = NewRequest().Body(body).Post(url).JsonResp(&resp)
	if err != nil {
		return false, err
	}

	return true, nil
}

//获取所在行业的信息
func (t *TempMsg) IndustryInfo() (string, error) {
	tk, err := token.Get()
	if err != nil {
		return "", err
	}

	url := HOST + "/cgi-bin/template/get_industry?access_token=" + tk
	resp, err := NewRequest().Post(url).String()
	if err != nil {
		return "", err
	}

	return resp, err
}

//获取模板列表
func (t *TempMsg) TempList() ([]TempInfo, error) {
	var resp []TempInfo
	tk, err := token.Get()
	if err != nil {
		fmt.Println(err)
	}

	url := HOST + "/cgi-bin/template/get_all_private_template?access_token=" + tk
	err = NewRequest().Post(url).JsonResp(&resp)
	if err != nil {
		fmt.Println(err)
	}

	return resp, err
}

//获取模板列表
func (t *TempMsg) GetTempId(no string) (TempMsgResp, error) {
	var resp TempMsgResp
	tk, err := token.Get()
	if err != nil {
		fmt.Println(err)
	}

	url := HOST + "/cgi-bin/template/api_add_template?access_token=" + tk
	err = NewRequest().Post(url).JsonResp(&resp)
	if err != nil {
		fmt.Println(err)
	}

	return resp, err
}

//删除模板
func (t *TempMsg) DelTemp(templateId string) (bool, error) {
	tk, err := token.Get()
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	body := `{ template_id:{{.templateId}} }`
	url := HOST + "/cgi-bin/template/del_private_template?access_token=" + tk
	var resp TempMsgResp
	err = NewRequest().Body(body).Post(url).JsonResp(&resp)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, nil
}

//发送
func (t *TempMsg) Send(data map[string]string) error {
	t.Config["data"] = data
	cfg, err := json.Marshal(t.Config)
	if err != nil {
		return err
	}

	var resp TempMsgResp
	err = NewRequest().Body(cfg).Get(t.url).JsonResp(&resp)
	if err != nil {
		return err
	}

	if resp.ErrCode != 0 {
		return errors.New("消息发送失败 msg:" + resp.ErrMsg)
	}

	return nil
}
