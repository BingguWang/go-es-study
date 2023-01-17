package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BingguWang/go-es-study/utils"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
	"testing"
)

func TestMget(t *testing.T) {
	req := esapi.MgetRequest{
		Index: "bing_index",
		Body: strings.NewReader(`{
		"docs" : [
      {
         "_index" : "bing_index",
         "_id" :    5
      },
      {
         "_index" : "bing_index",
         "_id" :    1,
         "_source": ["newFiled","name"]
      }
   ]
		}`),
	}
	res, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	utils.CheckEsSearchResErr(res)
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	fmt.Println(r)
}
