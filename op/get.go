package op

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/BingguWang/go-es-study/datasource"
	"github.com/BingguWang/go-es-study/utils"
	"github.com/BingguWang/go-es-study/vo"
	"github.com/elastic/go-elasticsearch/v8"
	"io"
	"log"
	"strings"
)

func GetAllIndices() error {
	log.Println("call GetAllIndices ...")
	client := datasource.GetESClient()
	// GET /_cat/indices/?v
	res, err := client.Cat.Indices(
		client.Cat.Indices.WithContext(context.Background()),
		client.Cat.Indices.WithPretty(),
		client.Cat.Indices.WithV(true),
	)
	if err != nil {
		log.Fatalf("[GetAllIndices] Error parsing the response body: %s", err)
		return err
	}
	defer res.Body.Close()
	log.Println(res)
	return nil
}

func GetIndexByIndexName(indexName string, v io.Reader) { // v是请求体
	client := datasource.GetESClient()
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(indexName),
		client.Search.WithBody(v),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
		client.Search.WithFrom(2), // 页码
		client.Search.WithSize(2), // 分页大小
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	utils.CheckEsSearchResErr(res)
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// 输出source
	var count int
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		count++
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}
	log.Println("docs count: ", count)
}

func MakeEsSearchRequestBody(query any) *bytes.Buffer {
	// 设置请求条件
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	return &buf
}

func GetClusterInfo() {
	log.Println("call GetClusterInfo ...")

	client := datasource.GetESClient()
	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
	var r map[string]interface{}
	// 反序列化
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))
}

func GetAllDocOfIndex(indexName string) {
	log.Println("call GetAllDocOfIndex ...")
	client := datasource.GetESClient()

	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(indexName),
		client.Search.WithPretty(),
		client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	var esRt vo.EsJsonResult
	//var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&esRt); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
		return
	}

	fmt.Println(utils.ToJson(esRt))
}
