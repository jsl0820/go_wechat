package wechat



type TicketResp struct {
	ErrCode   string `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

var ticket = &Ticket{expires:7200}

type Ticket struct {
	expires int
	Resp TicketResp 
}

//刷新
func (t *Ticket)Refresh(){
	t, err := token.Get()
	if err != nil {
		return nil, err
	}

	url := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?"
	url += "type=jsapi&access_token=" + t

	var s JsapiTicket
	err = NewRequest(&s).Get(url)
	if err != nil {
		return nil, error
	}

	return s, nil	
}

//获取票据
func (t *Ticket) Get() (string, error){

	if 	t.Resp == (TicketResp{}){
		t.Resp, err = t.Refresh()
		if err != nil {
			return nil , err
		}
		return  t.Resp.Ticket , nil 
	} 	

	return t.Resp.Ticket, nil
}


//定期清理
func (t *Ticket) Clear() {
	dur := time.Duration(t.expires) * time.Second
	for {
		<-time.After(dur)
		if t.Resp != (TicketResp{}) {
			t.Resp = TicketResp{}
		}
	}
}

type JsSdkConfig struct {
	Jt        string
	Url       string
	SingnData map[string]string
}

func (j *JsSdkConfig)Get() map[string]string {

	timeSamp := StampString()
	nonceStr := CreateNonceStr()

	j.SingnData = map[string]string{
		url:       j.Url,
		noncestr:  nonceStr,
		timestamp: timeSamp,
		jsapi_ticket : j.Ticket(),
	}

	signature := j.Sign() 

	return map[string]string{
		timesamp : timeSamp,
		noncestr : nonceStr,
		signature : signature,
	}
}


func init() {
	go ticket.Clear()		
}