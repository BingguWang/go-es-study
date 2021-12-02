package main

import (
	"context"
	"encoding/json"
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
	_, _, err = client.Ping(host).Do(context.Background())
	if err != nil {
		panic("连接es错误")
	}
	fmt.Println("连接es成功")
}

func main() {
	fmt.Println("----------------------------查 找 全 部---------------------------------")
	FindAll()

	fmt.Println("----------------------------id 查找,只返回一个结果---------------------------------")
	FindByIndexId()

	fmt.Println("----------------------------条件查找---------------------------------")
	FindByFiled()

}

func FindAll() {
	re, err := client.Search("user").Do(context.Background()) //searchService返回的是 *SearchResult
	if err != nil {
		panic("连接es出错：" + err.Error())
	}

	fmt.Printf("re.Hits.Hits：	%v \n", re.Hits.Hits) //真正的数据是在Hits的元素的source成员里
	for _, hit := range re.Hits.Hits {
		fmt.Printf("%v\n", hit.Source) // 可以看到source里的是JSON，需要解码
		fmt.Printf("%s\n", hit.Source) // 可以看到source里的是JSON，需要解码
		// JSON解码，hit的成员Source json.RawMessage  ，其实就是[]byte
		u := &entity.UserIndex{}
		if err := json.Unmarshal(hit.Source, u); err != nil { // Unmarshal必须要传指针
			panic(err)
		}
		fmt.Println(*u)
	}
	/**
		"_source": {
						"name": "wb",
						"age": 12,
						"married": false
	                }
	*/
	var userindex entity.UserIndex
	for i, item := range re.Each(reflect.TypeOf(userindex)) { // each里有对JSON进行解码, 其实也是对Source成员解码的
		fmt.Println(i)
		t := item.(entity.UserIndex)
		fmt.Printf("%#v\n", t)
	}
}

func FindByIndexId() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	res, err := client.Get().
		Index("user").
		Id("1").
		Do(context.Background()) // GetService的Do方法返回的是*GetResult
	if err != nil {
		fmt.Println(err)
		panic("按id查找失败")
	}
	// 数据是在GetResult的source成员中
	fmt.Printf("%s \n", res.Source)

	// 可以用JSON解码
	ui := entity.UserIndex{}
	if err := json.Unmarshal(res.Source, &ui); err != nil {
		panic("解析JSON出错")
	}
	fmt.Println(ui)

}

func FindByFiled() {
	q := elastic.NewQueryStringQuery("name:wb")
	res, err := client.Search("user").
		Query(q).
		Do(context.Background())

	//q := elastic.NewBoolQuery()
	//q.Must(elastic.NewQueryStringQuery("name:wb"), elastic.NewQueryStringQuery("age:12"))
	//res, err := client.Search("user").Query(q).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		panic("es 查找失败")
	}
	for _, hit := range res.Hits.Hits {
		a := entity.UserIndex{}
		json.Unmarshal(hit.Source, &a)
		fmt.Println(a)
	}

	// 或者下面这样使用searchResult自己封装的遍历结果的方法
	//var ui entity.UserIndex
	//for _, item := range res.Each(reflect.TypeOf(ui)) {
	//	fmt.Printf("%T \n",item)
	//	t := item.(entity.UserIndex)
	//	fmt.Println(t)
	//}
}
