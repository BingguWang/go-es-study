package utils

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
)

func CheckEsSearchResErr(res *esapi.Response) {
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
}
