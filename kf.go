package wechat

import (
	// "fmt"
	"errors"
)


type KfResp struct {
	ErrCode   int `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

type Kf struct {
	Account string `json:"kf_account"`
	Headimg string `json:"kf_headimgurl"`
	Id string `json:"kf_id"`
	Nick string `json:"kf_nick"`
	Wx string `json:"kf_wx"`
}

type KfList struct {
	List []Kf `json:"kf_list"` 
}


type KfAct struct {

}

// 添加
func (k *KfAct)Add(account, nickname string ) (bool,error){
	t, err := token.Get()
	if err != nil {
		return false, err
	}

	url := HOST + "/customservice/kfaccount/add?access_token=" + t
	body := `{kf_account:{{.account}}, nickname:{{.nickname}}}`

	var resp KfResp
	err = NewRequest().Body(body).Get(url).JsonResp(&resp)
	if err != nil {
		return false, err
	}

	if resp.ErrCode != 0 {
		return false, errors.New("请求失败错误码：" + resp.ErrMsg)
	}

	return true, nil 
}

//客服列表
func (k *KfAct)Get()(KfList, error){
	var resp KfList

	t, err := token.Get()
	if err != nil {
		return resp, err
	}

	url := HOST + "/cgi-bin/customservice/getonlinekflist?access_token" + t

	err = NewRequest().Get(url).JsonResp(&resp)
	if err != nil {
		return resp, err
	}
	
	return resp, nil 
}

//设置
func (k *KfAct)Set(account, nickname string) (bool,error){
	t, err := token.Get()
	if err != nil {
		return false, err
	}

	url := HOST + "ustomservice/kfaccount/update?access_token=" + t
	body := `{kf_account:{{.account}}, nickname:{{.nickname}}}`

	var resp KfResp
	err = NewRequest().Body(body).Get(url).JsonResp(&resp)
	if err != nil {
		return false, err
	}

	if resp.ErrCode != 0 {
		return false, errors.New("请求失败：" + resp.ErrMsg)
	}

	return true, nil 	
}

//删除
func (k *KfAct)Del(account string)(bool,error){
	t, err := token.Get()
	if err != nil {
		return false, err
	}

	url := HOST + "/customservice/kfaccount/del?access_token=" + t +"&kf_account=" + account
	var resp KfResp
	err = NewRequest().Get(url).JsonResp(&resp)
	if err != nil {
		return false, err
	}

	if resp.ErrCode != 0 {
		return false, errors.New("请求失败：" + resp.ErrMsg)
	}

	return true, nil 			
}

//绑定客服
func (k *KfAct)Bind(kfAccount, inviteWx string)(bool, error){
	t, err := token.Get()
	if err != nil {
		return false, err
	}

	url := HOST + "/customservice/kfaccount/inviteworker?access_token=" + t
	body := `{kf_account:{{.account}}, invite_wx:{{.inviteWx}}}`
	
	var resp KfResp
	err = NewRequest().Body(body).Get(url).JsonResp(&resp)
	if err != nil {
		return false, err
	}

	if resp.ErrCode != 0 {
		return false, errors.New("请求失败：" + resp.ErrMsg)
	}

	return true, nil 	
}

//上传头像
func (k *KfAct)UploadHeadimg(kfAccount, path string)error{
	t, err := token.Get()
	if err != nil {
		return err
	}

	url := HOST + "/customservice/kfaccount/uploadheadimg?access_token=" + t + "&kf_account=" + kfAccount

	req := NewRequest().Post(url)
	req.FormFile("filename", path)
	var resp KfResp
	err = req.JsonResp(&resp)
	if err != nil {
		return err 
	}

	if resp.ErrCode != 0 {
		return errors.New("操作失败：" + resp.ErrMsg)
	}

	return nil
}

