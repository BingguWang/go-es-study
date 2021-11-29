package main

import (
	"context"
	"fmt"
	elastic "github.com/olivere/elastic/v7"
	"github.com/wbing441282413/go-es-study/select/entity"
	"reflect"
)

var client *elastic.Client

const (
	host = "http://127.0.0.1:9200"
)

func init() {
	var err error
	client, err = elastic.NewClient(elastic.SetURL(host))
	if err != nil {
		panic("连接es出错")
	}
	fmt.Println("连接es成功")
}

func main() {
	//查找全部
	re, err := client.Search("user").Do(context.Background()) //searchService返回的是 *SearchResult
	if err != nil {
		panic("连接es出错：" + err.Error())
	}

	fmt.Printf("%v \n", re.Hits.Hits)
	for _, hit := range re.Hits.Hits {
		fmt.Printf("%#v\n", hit.Source) // 可以看到source里的是JSON，需要解码
	}

	var userindex entity.UserIndex
	for i, item := range re.Each(reflect.TypeOf(userindex)) {
		fmt.Println(i)
		t := item.(entity.UserIndex)
		fmt.Printf("%#v\n", t)
	}
}
