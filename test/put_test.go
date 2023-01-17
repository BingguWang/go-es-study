package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/BingguWang/go-es-study/datasource"
	"github.com/BingguWang/go-es-study/document"
	"github.com/BingguWang/go-es-study/utils"
	elastic8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
	"sync"
	"testing"
	"time"
)

/**
进行一些ES特性的测试
*/
var client *elastic8.Client

func init() {
	addrs := []string{
		"https://10.11.17.4:9200",
	}
	datasource.InitESClient("../datasource/ca.crt", addrs...)
	client = datasource.GetESClient()
}

// TestConcPutDoc 测试下并发put doc
func TestConcPutDoc(t *testing.T) {
	log.Println("call TestConcPutDoc ...")
	data := &document.UserDocument{
		Name:      "ooooooooooppppppppppppp",
		Age:       16,
		Married:   false,
		CreatedAt: time.Now(),
		About:     "",
	}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dt, _ := json.Marshal(data)
			// 这里的更新是全覆盖的，不能指定字段更新，要制定字段更新用client.update
			req := esapi.IndexRequest{
				Index:      "bing_index",
				DocumentID: "1", // id
				Body:       bytes.NewReader(dt),
				Refresh:    "true",
			}
			res, err := req.Do(context.Background(), client)
			if err != nil {
				log.Printf("[call TestConcPutDoc] Error getting response: %s", err)
				return
			}
			defer res.Body.Close()
			utils.CheckEsSearchResErr(res)
			if res.IsError() {
				log.Printf("[%s] Error indexing document : %v", res.Status(), res.String())
			} else {
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
			return
		}()
	}
	wg.Wait()
}
func TestUpdateByQuery(t *testing.T) {
	log.Println("call TestUpdate ...")
	res, err := client.UpdateByQuery(
		[]string{"bing_index"},
		client.UpdateByQuery.WithQuery("_id:1"), // query_string 查询 https://github.com/elastic/elasticsearch/blob/main/docs/reference/query-dsl/query-string-query.asciidoc
		client.UpdateByQuery.WithBody(
			strings.NewReader(`{
			  "script": {
				"source": "ctx._source.age += params.count",
				"lang": "painless",
				"params": {
				  "count": 1
				}
		  	}
		}`)),
		client.UpdateByQuery.WithRefresh(true),
	)
	defer res.Body.Close()
	fmt.Println(res)
	fmt.Println(err)
}

func TestUpdate(t *testing.T) {
	log.Println("call TestUpdate ...")
	// 只更新指定字段
	res, err := client.Update(
		"bing_index",
		"5",
		strings.NewReader(`{
			"doc":{
				"age":99
			}
		}`),
		client.Update.WithPretty(),
	)
	if err != nil {
		log.Printf("[call TestConcPutDoc] Error getting response: %s", err)
		return
	}
	defer res.Body.Close()
	utils.CheckEsSearchResErr(res)
	if res.IsError() {
		log.Printf("[%s] Error indexing document : %v", res.Status(), res.String())
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
	return
}

// 测试字段自增更新
func TestUpdateIncrease(t *testing.T) {
	log.Println("call TestUpdateIncrease ...")

	// 只更新指定字段
	res, err := client.Update(
		"bing_index",
		"5",
		// 自增1
		strings.NewReader(`{
	  "script": {
	    "source": "ctx._source.age += params.count",
	    "lang": "painless",
	    "params": {
	      "count": 1
	    }
	  }
	}`),
		client.Update.WithPretty(),
	)
	if err != nil {
		log.Printf("[call TestConcPutDoc] Error getting response: %s", err)
		return
	}
	defer res.Body.Close()
	utils.CheckEsSearchResErr(res)
	if res.IsError() {
		log.Printf("[%s] Error indexing document : %v", res.Status(), res.String())
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
	return
}
