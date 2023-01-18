package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BingguWang/go-es-study/datasource"
	"github.com/BingguWang/go-es-study/op"
	"github.com/BingguWang/go-es-study/utils"
	"github.com/BingguWang/go-es-study/vo"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
	"testing"
	"time"
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

func TestQueryDoc(t *testing.T) {
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"about": "rock",
	//		},
	//	},
	//}
	query := map[string]interface{}{}
	op.GetIndexByIndexName("bing_index", op.MakeEsSearchRequestBody(query))
}

func TestGetAllDocOfIndex(t *testing.T) {
	op.GetAllDocOfIndex("bing_index")
}

func TestScroll(t *testing.T) {
	client := datasource.GetESClient()

	query := map[string]interface{}{}
	res, err := client.Search(
		client.Search.WithSize(5),
		client.Search.WithScroll(2*time.Minute),
		client.Search.WithBody(op.MakeEsSearchRequestBody(query)),
		client.Search.WithPretty(),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithIndex("bing_index"),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	var r vo.EsJsonResult
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
		return
	}
	fmt.Println(len(r.Hits.Hits))
	fmt.Println(utils.ToJson(r.Hits.Hits))
	for {
		response, err := client.Scroll(
			client.Scroll.WithScrollID(r.ScrollId),
			client.Scroll.WithScroll(time.Minute),
		)
		if err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		}
		defer response.Body.Close()
		//var rs map[string]interface{}
		var rs vo.EsJsonResult
		if err := json.NewDecoder(response.Body).Decode(&rs); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
			return
		}
		fmt.Println(len(rs.Hits.Hits))
		fmt.Println(utils.ToJson(rs.Hits.Hits))
		if len(rs.Hits.Hits) == 0 {
			return
		}
	}
}
