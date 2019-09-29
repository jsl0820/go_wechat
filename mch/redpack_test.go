package mch

import (
	"testing"
)

func TestSendRedPack(t *testing.T) {

	info := map[string]string{
		"mch_billno":   "",
		"send_name":    "",
		"re_openid":    "",
		"total_amount": "",
		"total_num":    "",
		"wishing":      "",
		"client_ip":    "",
		"act_name":     "",
		"remark":       "remark",
	}

	redPack := &RedPack{Info: info}
	redPack.SendRedPack()
}

func TestRecord(t *testing.T) {
	info := map[string]string{
		"bill_type":  "",
		"mch_billno": "",
	}
	redPack := &RedPack{Info: info}
	redPack.Record(1)
}
