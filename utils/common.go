package utils

import (
	"encoding/json"
	"time"
)

func ToJson(i interface{}) string {
	marshal, _ := json.Marshal(i)
	return string(marshal)
}

func GetTimePtr(t time.Time) *time.Time {
	cstSh, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		cstSh = time.FixedZone("CST", 8*3600)
	}
	in := t.In(cstSh)
	return &in
}
