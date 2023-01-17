package utils

import (
	"encoding/json"
	"github.com/BingguWang/go-es-study/document"
	"math/rand"
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

// GetRandomString 随机生成字符串
func GetRandomString(l int) string {
	str := "abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := make([]byte, l)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
	}
	return string(result)
}

func GetRandomMsg(l int) string {
	str := "abcdefghijk lmnopqrstuvwxyz        !."
	bytes := []byte(str)
	result := make([]byte, l)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
	}
	return string(result)
}

func GetRandomUserDoc() *document.UserDocument {
	var married bool
	if (time.Now().Second() % 2) == 0 {
		married = true
	}
	r := &document.UserDocument{
		Name:      GetRandomString(6),
		Age:       rand.Intn(100),
		Married:   married,
		CreatedAt: time.Now().UTC(),
		About:     GetRandomMsg(50),
	}
	return r
}
