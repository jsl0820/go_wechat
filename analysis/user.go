package analysis

type UserSummary struct {
	RefDate string `json:"ref_date"`
	UserSource int `json:"user_source"`
	NewUser int `json:"new_user"`
	CancelUser string `json:"cancel_user"`
}

type UserCumulate struct{
	RefDate string `json:"ref_date"`
	UserSource int `json:"user_source"`
	NewUser int `json:"new_user"`
	CancelUser string `json:"cancel_user"`
} 


func (a *Analysis)User(respType int)(*Analysis){
	return a.setToken()
}


func (a *Analysis)Summary()(SummaryResp, error){
	url := HOST + "datacube/getusersummary?access_token=" + a.token
	body, err := json.Marshal(a) 
	var Resp SummaryResp
	if err != nil {
		return Resp, err 
	}

	req := NewRequest().Body(body).Post(url)
	if err := req.JsonResp(&Resp); err != nil {
		return Resp, err 
	} else {
		return Resp, nil
	}
}

func (a *Analysis)Cumulate()(CumulateResp, error){
	url := HOST + "datacube/getusersummary?access_token=" + a.token
	body, err := json.Marshal(a) 
	var Resp CumulateResp
	if err != nil {
		return Resp, err 
	}

	req := NewRequest().Body(body).Post(url)
	if err := req.JsonResp(&Resp); err != nil {
		return Resp, err 
	} else {
		return Resp, nil
	}
}
