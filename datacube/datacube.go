package datacube

import (
	"log"
	"strings"

	wx "github.com/wechat"
)

type Datacube struct {
	url       string
	token     string
	respType  int
	EndDate   string `json:"end_date"`
	BeginDate string `json:"begin_date"`
}

func (d *Datacube) url(url string) string {
	if token, err := wx.TokenInstance.Get(); err != nil {
		panic(err)
		log.Println(err)
	}

	return strings.Replace(url, "{{TOKEN}}", token, -1)
}
