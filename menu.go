package wechat

import(
	"encoding/json"
	"log"
	"errors"
)

type MenueResp struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

type Menu struct {

} 

//查询
func (m *Menu) Query(cat int)(string,error){
	t, err := token.Get()
	if err != nil {
		return "", err
	}

	var url string 
	if cat == 0 {
		url = HOST + "/get_current_selfmenu_info?access_token=" + t 
	}
	if cat == 1{
		url = HOST + "/get_current_selfmenu_info?access_token=" + t 
	}
	
	resp, err := NewRequest().Get(url).String()

	if err != nil {
		return "", err
	}

	return resp, nil 	
}


//创建
func (m *Menu)Create(menu map[string]interface{})error{
	jsonStr, err := json.Marshal(menu)
	if err != nil {
		log.Println(err)
	}

	t, err := token.Get()
	if err != nil {
		return  err
	}

	url :=  HOST + "/cgi-bin/menu/create?access_token=" + t
	var resp MenueResp
	err = NewRequest().Body(jsonStr).Post(url).JsonResp(&resp)
	if err != nil {
		return err
	}

	if resp.ErrCode != 0 {
		return errors.New("发生错误" +  resp.ErrMsg)
	}

	return nil
}

//获取菜单
func (m *Menu)Get(cat int)( string, error){
	t, err := token.Get()
	if err != nil {
		return  "", err
	}

	var url string
	if cat == 1{
		url =  HOST + "/cgi-bin/menu/get?access_token=" + t
	}
	if cat == 0{
		url = HOST + "/cgi-bin/menu/addconditional?access_token=" + t
	}

	resp, err := NewRequest().Get(url).String()
	if err != nil {
		return "", err
	}

	return resp, nil
}

//删除全部菜单
func(m *Menu) Del()(bool, error){
	t, err := token.Get()
	if err != nil {
		return  false,  err
	}

	url :=  HOST + "/cgi-bin/menu/delete?access_token=" + t
	var resp MenueResp
	err = NewRequest().Get(url).JsonResp(&resp)

	if err != nil {
		return false, err
	}

	return false, nil
}
