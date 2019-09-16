package datacube

type Datacube struct {
	url       string
	token     string
	respType  int
	EndDate   string `json:"end_date"`
	BeginDate string `json:"begin_date"`
}
