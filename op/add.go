package op

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/BingguWang/go-es-study/datasource"
	"github.com/BingguWang/go-es-study/utils"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"time"
)

// 新增一个document
func IndexDoc(indexName string, id string, version int, data interface{}) {
	log.Println("call IndexDoc ...")

	client := datasource.GetESClient()
	dt, _ := json.Marshal(data)
	req := esapi.IndexRequest{
		Index: indexName,
		//DocumentID: "", // id
		Body:    bytes.NewReader(dt),
		Refresh: "true",
		Timeout: 5 * time.Second,
		//Version: nil, // 版本号，执行的时候先对比版本号，不一直就错误
	}
	if id != "" {
		req.DocumentID = id
	}
	if version > 0 {
		req.Version = &version
	}
	res, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("[call IndexDoc] Error getting response: %s", err)
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
}
