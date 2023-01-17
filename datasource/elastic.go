package datasource

import (
	"crypto/x509"
	elastic8 "github.com/elastic/go-elasticsearch/v8"
	"io/ioutil"
	"log"
	"net/http"
)

var client *elastic8.Client

func GetESClient() *elastic8.Client {
	return client
}

func InitESClient(crtPath string, addr ...string) {
	var err error
	ca, err := ioutil.ReadFile(crtPath)
	if err != nil {
		panic(err)
	}
	// --> Clone the default HTTP transport
	tp := http.DefaultTransport.(*http.Transport).Clone()

	// --> Initialize the set of root certificate authorities
	if tp.TLSClientConfig.RootCAs, err = x509.SystemCertPool(); err != nil {
		log.Fatalf("ERROR: Problem adding system CA: %s", err)
	}

	// --> Add the custom certificate authority
	if ok := tp.TLSClientConfig.RootCAs.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("ERROR: Problem adding CA from file ")
	}
	cfg := elastic8.Config{
		Addresses: addr,
		Username:  "elastic",
		Password:  "wb123456",
		// --> Pass the transport to the client
		Transport: tp,
	}
	newClient, err := elastic8.NewClient(cfg)
	if err != nil {
		log.Fatalln("new client error : ", err)
	}
	// test if new client successfully
	res, err := newClient.Info()
	if err != nil {
		log.Fatalf("ERROR: Unable to get response: %s", err)
	}
	log.Println(res)

	log.Println("new client successfully on addr : ", addr)
	client = newClient

}
