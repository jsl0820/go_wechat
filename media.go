package wechat

import (
	"fmt"
	"errors"
)

type CountResp struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
	CoiceCount int `json:"voice_count"`
	VideoCount int `json:"video_count"`
	ImageCount int `json:"image_count"`
	NewsCount  int `json:"news_count"` 
}

type MediaBaseResp struct{
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

type UploadResp struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
	Type string `json:"type"`
	MediaId string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

type Media struct {
	
}

//上传图片
func(m *Media)UploadImg(filename string)(string, error){
	tk, err := token.Get()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	url := HOST + "/cgi-bin/media/upload?access_token="+ tk + "&type=image" 
	var resp UploadResp
	err = NewRequest().Upload(filename).Post(url).JsonResp(&resp)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return resp.MediaId, nil 
}

//上传声音
func(m *Media)UploadVoice()(string, error){
	tk, err := token.Get()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	url := HOST + "/cgi-bin/media/upload?access_token="+ tk + "&type=voice" 
	var resp UploadResp
	err = NewRequest().Upload(filename).Post(url).JsonResp(&resp)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return resp.MediaId, nil 
}

//上传视频
func(m *Media)UploadVedio(filename string)(string, error){
	tk, err := token.Get()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	url := HOST + "/cgi-bin/media/upload?access_token="+ tk + "&type=video" 
	var resp UploadResp
	err = NewRequest().Upload(filename).Post(url).JsonResp(&resp)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return resp.MediaId, nil 
}

//上传缩略图
func (m *Media)UploadThumb(filename string)(string, error){
	tk, err := token.Get()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	url := HOST + "/cgi-bin/media/upload?access_token="+ tk + "&type=video" 
	var resp UploadResp
	err = NewRequest().Upload(filename).Post(url).JsonResp(&resp)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return resp.MediaId, nil 
}

//下载素材
func (m *Media)Download(id string){

}

//删除
func (m *Media)Del(id string)(bool, error){
	tk, err := token.Get()
	if err !=nil {
		return false, err
	}

	body := `{media_id:{{.id}} }`
	url := HOST + "/cgi-bin/material/del_material?access_token=" + tk
	var resp MediaBaseResp
	err = NewRequest().Body(body).Post(url).JsonResp(&resp)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if resp.ErrCode != 0 {
		return false, errors.New("操作失败：", resp.ErrMsg)
	}

	return true, nil
}


//统计
func (m *Media)Count()(CountResp, error){
	tk, err := token.Get()
	if err != nil {
		return CountResp{}, err
	}

	url := HOST + "/cgi-bin/material/get_materialcount?access_token=" + tk
	var resp CountResp
	err = NewRequest().Get(url).JsonResp(&resp)
	if err != nil {
		return resp, nil 
	}

	return resp, nil
}

