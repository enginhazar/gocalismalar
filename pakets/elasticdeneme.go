package pakets

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	_ "github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"os"
	"strings"
)

type ElasticSearch struct {
	client *elasticsearch.Client
	index  string
	alias  string
}

func NewElasticClient() (*ElasticSearch, error) {
	fmt.Printf("new client")
	cfg := elasticsearch.Config{

		Addresses: []string{"http://localhost:9200/"},
		Username:  "elastic",
		Password:  "123456",
		Logger:    &elastictransport.ColorLogger{Output: os.Stdout},
		//Logger:    &estransport.ColorLogger{os.Stdout, true, true},
	}
	fmt.Println("config")
	eclient, _err := elasticsearch.NewClient(cfg)

	if _err != nil {
		log.Fatalln("err")
		return nil, _err
	}

	return &ElasticSearch{client: eclient}, nil
}
func (e *ElasticSearch) InsertDocument(s Sicil) error {

	data, _err := json.Marshal(s)

	//	e.client.Info() //// Çalışıyor

	if _err != nil {
		return _err
	}

	res, err := e.client.Index(
		e.alias,                         // Dokümanın ekleneceği indeks
		strings.NewReader(string(data)), // Doküman verisi
		e.client.Index.WithContext(context.Background()),
	)

	//req := esapi.IndexRequest{Index: e.alias, DocumentID: string(s.Sicilno), Body: strings.NewReader(string(data))}

	//res, err := req.Do(context.Background(), e.client)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("insert: request: %w", err)
	}

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("insert: request: %w", err)
	}
	println(res.Body)
	res.Body.Close()

	if res.StatusCode == 409 {
		fmt.Println("status code 409")
		return fmt.Errorf("insert: request:409")
		return nil
	}

	//req := esapi.CreateRequest{
	//	Index:      e.alias,
	//	DocumentID: string(s.Sicilno),
	//	Body:       bytes.NewReader(data),
	//}
	////fmt.Println(req.)
	//
	//res, err := req.Do(context.Background(), e.client)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return fmt.Errorf("insert: request: %w", err)
	//}
	//defer res.Body.Close()
	//
	//if res.StatusCode == 409 {
	//	fmt.Println("status code 409")
	//	return fmt.Errorf("insert: request:409")
	//	return nil
	//}
	//
	//if res.IsError() {
	//	return fmt.Errorf("insert: response: %s", res.String())
	//}
	//fmt.Println("doküman kayıt edildi")

	return nil
}

func (e *ElasticSearch) CreateIndex(index string) error {
	e.index = index
	e.alias = index + "_alias"

	res, _err := e.client.Indices.Exists([]string{e.index})

	if _err != nil {
		log.Fatalln(_err)
		return _err
	}

	if res.StatusCode == 200 {
		return nil
	}

	if res.StatusCode == 400 {
		return fmt.Errorf("error Indices : %s", res.String())
	}

	res, _err = e.client.Indices.Create(e.index)
	if _err != nil {
		return fmt.Errorf("error create : %s", res.String())
	}

	if res.IsError() {
		return fmt.Errorf("İndex Create error %s", res.String())
	}

	res, _err = e.client.Indices.PutAlias([]string{e.index}, e.alias)

	if _err != nil {
		log.Fatalln(_err)
		return _err
	}
	if res.IsError() {
		return fmt.Errorf("İndex Alias error %s", res.String())
	}
	fmt.Println("Index oluşturuldu")
	return nil
}
