package analysis

import (
	"encoding/json"
	"log"

	wx "github.com/go_wechat"
)

const HOST = wx.HOST
 
type Analysis struct {
	BeginDate string  `json:"begin_date"`
	EndDate   string  `json:"end_date"`
	token	  string 	
	respType  int
}

func (a *Analysis)setToken()*Analysis{
	if a.token, err := wx.TokenInstance.Get(); err != nil {
		panic(err)
		log.Println(err)
	} 

	return a
}

type SummaryResp struct {
	List []UserSummary `json:"list"`
	data []byte
}

type CumulateResp struct {
	List []UserCumulate `json:"list"`
}


func (r *CumulateResp)UnMarshaled(b []byte)(CumulateResp, error){
	if err := json.Unmarshal(b, r); err != nil {
		return  *r, err
	} else {
		return *r, nil
	}
}

