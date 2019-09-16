package datacube

const INTERFACE_SUMMARY = "/datacube/getinterfacesummary?access_token={{TOKEN}}"
const INTERFACE_SUMMARY_HOUR = "/datacube/getinterfacesummaryhour?access_token={{TOKEN}}"

type InterfaceResp struct {
	RefDate       string
	CallbackCount uint
	FailCount     uint
	TotalTimeCost uint
	MaxTimeCost   uint
}

// type InterfaceRespList InterfaceResp{}

type Interface struct {
	Datacube
}

func (i *Interface) Summary() *Analysis {
	url := i.url(INTERFACE_SUMMARY)
	req := NewRequest().Body(body).Post(url)

	var resp InterfaceResp
	if err := req.JsonResp(&resp); err != nil {
		return resp, err
	}

	return resp, nil
}
