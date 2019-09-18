package datacube

const MSG_SG = "/datacube/getupstreammsg?access_token={{TOKEN}}"
const MSG_HOUR = "/datacube/getupstreammsghour?access_token={{TOKEN}}"
const MSG_WEEK = "/datacube/getupstreammsgweek?access_token={{TOKEN}}"
const MSG_MONTH = "/datacube/getupstreammsgmonth?access_token={{TOKEN}}"
const MSG_DIST = "/datacube/getupstreammsgdist?access_token={{TOKEN}}"
const MSG_DIST_WEEK = "/datacube/getupstreammsgdistweek?access_token={{TOKEN}}"
const MSG_DIST_MONTH = "/datacube/getupstreammsgdistmonth?access_token={{TOKEN}}"

type MsgResp struct {
	RefDate       string
	RefHour       int
	CallbackCount uint
	FailCount     uint
	TotalTimeCost uint
	MaxTimeCost   uint
}

type MsgResp struct {
	List []MsgResp `json:"list"`
}

type Message struct {
	Datacube
}

func (m *MsgResp) Msg() MsgResp {

}

func (m *Message) Hour() {

}

func (m *Message) Week() {

}

func (m *Message) Month() {

}

func (m *Message) Dist() {

}

func (m *Message) Distweek() {

}

func (m *Message) Distmonth() {
	
}
