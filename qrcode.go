package wechat

type QrcResp struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds string `json:"expire_seconds"`
	Url           string `json:"url"`
}

type Qrcode struct {
	Resp    QrcResp
	token string
	JsonString string
}

// 参数说明
// 生成临时二维码
// 1.带参数的二维码 
// {"expire_seconds": 604800, "action_name": "QR_SCENE", "action_info": {"scene": {"scene_id": 123}}}
// 2.字符串形式的二维码参数
// {"expire_seconds": 604800, "action_name": "QR_STR_SCENE", "action_info": {"scene": {"scene_str": "test"}}}
// 永久二维码
// 1.带参数的二维码
// {"action_name": "QR_LIMIT_SCENE", "action_info": {"scene": {"scene_id": 123}}}
// 2.字符串形式的二维码参数
// {"action_name": "QR_LIMIT_STR_SCENE", "action_info": {"scene": {"scene_str": "test"}}}
// 

func QrCode(jsonString string) Qrcode {
	return &Qrcode{
		JsonString:jsonString,
	}
}

func (q *Qrcode)Resp() (QrcResp, error) {
	var resp QrcResp
	url := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + q.token
	err := NewRequest(&resp).JsonPost(q.JsonString, url)
	if err != nil {
		return nil , err
	} 
	return resp, nil
} 

func (q *Qrcode) DownLoad(savePath , fileName string) {
	resp, err  :=  q.Resp()
	imgUrl := "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + resp.Ticket 

}
