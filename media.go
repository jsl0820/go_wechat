package wechat

import (
	"log"
	"errors"
)

type CountResp struct {
	Url string `json:"url"`
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
	CoiceCount int `json:"voice_count"`
	VideoCount int `json:"video_count"`
	ImageCount int `json:"image_count"`
	NewsCount  int `json:"news_count"` 
}

type MediaResp struct{
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
	CoiceCount int `json:"voice_count"`
	VideoCount int `json:"video_count"`
	ImageCount int `json:"image_count"`
	NewsCount  int `json:"news_count"` 
	Type string `json:"type"`
	MediaId string `json:"media_id"`
	CreatedAt string `json:"created_at"`
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

//上传临时素材
func (m *Media)Upload(fileType , filename string)(UploadResp, error){
	var resp UploadResp
	tk, err := token.Get()
	if err != nil {
		log.Println(err)
		return resp, err
	}

	url := HOST + "/cgi-bin/media/upload?access_token="+ tk + "&type=" + fileType 
	req := NewRequest().Post(url)
	req.FormFile("media", filename)
	err = req.Form().JsonResp(&resp)
	if err != nil {
		return resp, err
	}

	if resp.ErrMsg != ""{
		return resp, errors.New("请求出错" + resp.ErrMsg)
	}

	return resp, nil
}


//上传图片
func(m *Media)UploadImg(filename string)(string, error){
	resp, err := m.Upload("image", filename) 
	if err != nil {
		return "", err 
	}

	return resp.MediaId, nil 
}

//上传声音
// func(m *Media)UploadVoice()(string, error){
// 	tk, err := token.Get()
// 	if err != nil {
// 		log.Println(err)
// 		return "", err
// 	}

// 	url := HOST + "/cgi-bin/media/upload?access_token="+ tk + "&type=voice" 
// 	var resp UploadResp
// 	err = NewRequest().Upload(filename).Post(url).JsonResp(&resp)
// 	if err != nil {
// 		fmt.Println(err)
// 		return "", err
// 	}

// 	return resp.MediaId, nil 
// }

//上传视频
// func(m *Media)UploadVedio(filename string)(string, error){
// 	tk, err := token.Get()
// 	if err != nil {
// 		fmt.Println(err)
// 		return "", err
// 	}

// 	url := HOST + "/cgi-bin/media/upload?access_token="+ tk + "&type=video" 
// 	var resp UploadResp
// 	err = NewRequest().Upload(filename).Post(url).JsonResp(&resp)
// 	if err != nil {
// 		log.Println(err)
// 		return "", err
// 	}

// 	return resp.MediaId, nil 
// }

//上传缩略图
// func (m *Media)UploadThumb(filename string)(string, error){
// 	tk, err := token.Get()
// 	if err != nil {
// 		log.Println(err)
// 		return "", err
// 	}

// 	url := HOST + "/cgi-bin/media/upload?access_token="+ tk + "&type=video" 
// 	var resp UploadResp
// 	err = NewRequest().Upload(filename).Post(url).JsonResp(&resp)
// 	if err != nil {
// 		log.Println(err)
// 		return "", err
// 	}

// 	return resp.MediaId, nil 
// }

// //下载素材
// func (m *Media)Download(id, path string) error {
// 	tk, err := token.Get()
// 	if err != nil {
// 		return err
// 	}

// 	var resp MediaResp
// 	url := HOST + "/cgi-bin/media/get?access_token=" + tk + "&media_id=" + id
// 	err = NewRequest().Get(url).SaveTo(path)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }


//获取视频
// func (m *Media)GetVedio(id string)(string, error){
// 	tk, err := token.Get()
// 	if err != nil {
// 		return "", err
// 	}

// 	var resp MediaResp
// 	url := HOST + "/cgi-bin/media/get?access_token=" + tk + "&media_id=" + id
// 	err := NewRequest().Get(url).Resp(&resp)
// }

// //删除
// func (m *Media)Del(id string)(bool, error){
// 	tk, err := token.Get()
// 	if err !=nil {
// 		return false, err
// 	}

// 	body := `{media_id:{{.id}} }`
// 	url := HOST + "/cgi-bin/material/del_material?access_token=" + tk
// 	var resp MediaBaseResp
// 	err = NewRequest().Body(body).Post(url).JsonResp(&resp)
// 	if err != nil {
// 		fmt.Println(err)
// 		return false, err
// 	}

// 	if resp.ErrCode != 0 {
// 		return false, errors.New("操作失败：", resp.ErrMsg)
// 	}

// 	return true, nil
// }


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

//获取素材列表
//t:类型 参数：'image', 'vedio', 'voice', 'news' 
//off:偏移量 count:返回的数量
func (m *Media) List(t, off, count string){

}

type Article struct {
	Title string `json:"title"`
	ThumbMediaId string `json:"thumb_media_id"`
	Author string `json:"author"` 
	Digest string `json:"digest"`
	ShowCoverPic int `json:"show_cover_pic"`
	Content string `json:"content"`
	ContentSourceUrl string `json:"content_source_url"`
	NeedOpenComment int  `json:"need_open_comment"`
	OnlyFansCanComment int `json:"only_fans_can_comment"`
}

//上传图文
// func(a *Article)Upload()(string, error){
// 	tk, err := token.Get()
// 	if err !=nil {
// 		return "", err
// 	}

// 	var resp &MediaResp
// 	body, err := json.Marshal(*a)
// 	url := HOST + "/cgi-bin/material/add_news?access_token=" + ACCESS_TOKEN
// 	err := NewRequest().Body(body).Post(url).Resp(&)
// 	if err != nil {
// 		return "", err 
// 	}

// 	return resp.MediaId, nil
// } 

// //上传图文消息的图片
// func (a *Article)UploadImg(path string)(string, error){
// 	tk, err := token.Get()
// 	if err != nil {
// 		return "", err
// 	}

// 	url := HOST + "/cgi-bin/media/uploadimg?access_token=" + tk
// 	NewRequest().Upload(path).Post()
// }

//获取
func (a *Article)Get(id string){

}





//图片
type Image struct {

}

//上传
func (i *Image)Upload(){

}

//下载
func (i *Image)Download(){

}

//删除
func (i *Image)Del(){

}


//视频
type Voice struct {

}



//缩略图
type Thumb struct{

}
