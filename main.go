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
	host   = flag.String("host", "localhost", "")
	port   = flag.String("port", "9200", "")
)

func main() {
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)
	fmt.Println(addr)

	datasource.InitESClient("http://" + addr)

	//op.GetClusterInfo()

	op.GetAllIndices()
}

//
//func FindAll() {
//	re, err := client.Search("user").Do(context.Background()) //searchService返回的是 *SearchResult
//	if err != nil {
//		panic("连接es出错：" + err.Error())
//	}
//
//	fmt.Printf("re.Hits.Hits：	%v \n", re.Hits.Hits) //真正的数据是在Hits的元素的source成员里
//	for _, hit := range re.Hits.Hits {
//		fmt.Printf("%v\n", hit.Source) // 可以看到source里的是JSON，需要解码
//		fmt.Printf("%s\n", hit.Source) // 可以看到source里的是JSON，需要解码
//		// JSON解码，hit的成员Source json.RawMessage  ，其实就是[]byte
//		u := &document.UserDocument{}
//		if err := json.Unmarshal(hit.Source, u); err != nil { // Unmarshal必须要传指针
//			panic(err)
//		}
//		fmt.Println(*u)
//	}
//	/**
//	  	"_source": {
//	  					"name": "wb",
//	  					"age": 12,
//	  					"married": false
//	                  }
//	*/
//	var userindex document.UserDocument
//	for i, item := range re.Each(reflect.TypeOf(userindex)) { // each里有对JSON进行解码, 其实也是对Source成员解码的
//		fmt.Println(i)
//		t := item.(document.UserDocument)
//		fmt.Printf("%#v\n", t)
//	}
//}
//
//func FindByIndexId() {
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println(err)
//		}
//	}()
//	res, err := client.Get().
//		Index("user").
//		Id("1").
//		Do(context.Background()) // GetService的Do方法返回的是*GetResult
//	if err != nil {
//		fmt.Println(err)
//		panic("按id查找失败")
//	}
//	// 数据是在GetResult的source成员中
//	fmt.Printf("%s \n", res.Source)
//
//	// 可以用JSON解码
//	ui := document.UserDocument{}
//	if err := json.Unmarshal(res.Source, &ui); err != nil {
//		panic("解析JSON出错")
//	}
//	fmt.Println(ui)
//
//}
//
//func FindByFiled() {
//	q := elastic.NewQueryStringQuery("name:wb")
//	res, err := client.Search("user").
//		Query(q).
//		Do(context.Background())
//
//	//q := elastic.NewBoolQuery()
//	//q.Must(elastic.NewQueryStringQuery("name:wb"), elastic.NewQueryStringQuery("age:12"))
//	//res, err := client.Search("user").Query(q).Do(context.Background())
//
//	if err != nil {
//		fmt.Println(err)
//		panic("es 查找失败")
//	}
//	for _, hit := range res.Hits.Hits {
//		a := document.UserDocument{}
//		json.Unmarshal(hit.Source, &a)
//		fmt.Println(a)
//	}
//
//	// 或者下面这样使用searchResult自己封装的遍历结果的方法
//	//var ui document.UserDocument
//	//for _, item := range res.Each(reflect.TypeOf(ui)) {
//	//	fmt.Printf("%T \n",item)
//	//	t := item.(document.UserDocument)
//	//	fmt.Println(t)
//	//}
//}
