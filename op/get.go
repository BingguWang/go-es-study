package op

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BingguWang/go-es-study/datasource"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"strings"
)

func GetAllIndices() error {
	log.Println("call GetAllIndices ...")
	client := datasource.GetESClient()
	// todo 未生效
	res, err := client.Cat.Indices(
		client.Cat.Indices.WithContext(context.Background()),
	)
	if err != nil {
		return err
	}
	var r map[string]interface{}
	json.NewDecoder(res.Body).Decode(&r)
	fmt.Println(r)
	return nil
}

func GetIndexByIndexName() {

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
