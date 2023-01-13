package datasource

import (
	elastic8 "github.com/elastic/go-elasticsearch/v8"
	"log"
)

var client *elastic8.Client

func GetESClient() *elastic8.Client {
	return client
}

func InitESClient(addr ...string) {
	cfg := elastic8.Config{
		Addresses: addr,
		Username:  "elastic",
		Password:  "123456",
	}
	newClient, err := elastic8.NewClient(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("new client successfully on addr : ", addr)
	client = newClient

}
