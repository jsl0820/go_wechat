package wechat

import (

	"encoding/json"
	"log"
)

type resp interface {
	Byte([]byte)[]byte
	UnMarshaled([]byte)(resp, error)
}


type Analysis struct {
	BeginDate string  `json:"begin_date"`
	EndDate   string  `json:"end_date"`
	token	  string 	
	respType  int
}

func (a *Analysis)setToken()*Analysis{
	if token, err := token.Get(); err != nil {
		panic(err)
		log.Println(err)
	} else {
		a.token = token
	}

	return a
}

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


type SummaryResp struct {
	List []UserSummary
}

func (r *SummaryResp)Bytes(b []byte)[]byte{
	return b
}

func (r *SummaryResp)UnMarshaled(b []byte)(SummaryResp, error){
	if err := json.Unmarshal(b, r); err != nil {
		return  *r, err
	} else {
		return *r, nil
	}
}

type CumulateResp struct {
	List []UserCumulate
}

func (r *CumulateResp)Bytes(b []byte)[]byte{
	return b
}

func (r *CumulateResp)UnMarshaled(b []byte)(CumulateResp, error){
	if err := json.Unmarshal(b, r); err != nil {
		return  *r, err
	} else {
		return *r, nil
	}
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




func (a *Analysis)Article()(*Analysis){
	return a.setToken()
}

