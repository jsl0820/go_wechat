package wechat


type UserData struct {
	RefDate string `json:"ref_date"`
	UserSource int `json:"user_source"`
	NewUser int `json:"new_user"`
	CancelUser string `json:"cancel_user"`
}

type UserDataResp struct {
	List []UserData
}


type Analysis struct {
	
}

func (a *Analysis)reqBody()string{

}

func (a *Analysis)User(begin, end string)(UserDataResp, error){
	t, err := token.Get()
	var resp UserDataResp
	if err != nil {
		return resp, err
	}

	url := HOST + "ustomservice/kfaccount/update?access_token=" + t
	body := `{kf_account:{{.account}}, nickname:{{.nickname}}}`


}