package main

import (
	"flag"
	"fmt"
	"github.com/BingguWang/go-es-study/datasource"
	"github.com/BingguWang/go-es-study/op"
	elastic8 "github.com/elastic/go-elasticsearch/v8"
	"net"
)

var (
	client *elastic8.Client
	//host   = flag.String("host", "localhost", "")
	host = flag.String("host", "139.9.221.92", "")
	port = flag.String("port", "9200", "")
)

func main() {
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)
	fmt.Println(addr)
	addrs := []string{
		"https://" + addr,
	}
	datasource.InitESClient("./datasource/ca.crt", addrs...)

	//op.IndexDoc("bing_index", "", 0, utils.GetRandomUserDoc())

	//op.GetClusterInfo()
	//op.GetAllIndices()

	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"about": "rock",
	//		},
	//	},
	//}
	query := map[string]interface{}{}
	op.GetIndexByIndexName("bing_index", op.MakeEsSearchRequestBody(query))

	op.GetAllDocOfIndex("bing_index")
}
