package wechat

import (
	// "fmt"
	"errors"
)


type KfResp struct {
	ErrCode   int `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
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



// //设置
// func (k *KfAct)Set(account, nickname string) (bool,error){
// 	t, err := token.Get()
// 	if err != nil {
// 		return false, err
// 	}

// 	url := HOST + "ustomservice/kfaccount/update?access_token=" + t
// 	body := `{kf_account:{{.account}}, nickname:{{.nickname}}}`

// 	var resp KfResp
// 	err = NewRequest().Body(body).Get(url).JsonResp(&resp)
// 	if err != nil {
// 		return false, err
// 	}

// 	if resp.ErrCode != 0 {
// 		return false, errors.New("请求失败错误码：" + resp.ErrCode)
// 	}

// 	return true, nil 	
// }

// //删除
// func (k *KfAct)Del(account string)(bool,error){
// 	t, err := token.Get()
// 	if err != nil {
// 		return false, err
// 	}

// 	url := HOST + "/customservice/kfaccount/del?access_token=" + t +"&kf_account=" + account

// 	var resp KfResp
// 	err = NewRequest().Get(url).JsonResp(&resp)
// 	if err != nil {
// 		return false, err
// 	}

// 	if resp.ErrCode != 0 {
// 		return false, errors.New("请求失败错误码：" + resp.ErrCode)
// 	}

// 	return true, nil 			
// }

//上传头像
// 这里先不写
func (k *KfAct)UploadHeadimg(){

}

