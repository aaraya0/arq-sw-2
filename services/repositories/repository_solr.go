package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aaraya0/arq-software/arq-sw-2/dtos"

	"fmt"

	e "github.com/aaraya0/arq-software/arq-sw-2/utils/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	logger "github.com/sirupsen/logrus"
	"github.com/stevenferrer/solr-go"
)

type Publi struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	City        string `json:"city"`
	State       string `json:"state"`
	Image       string `json:"image"`
	Seller      string `json:"seller"`
}
type SolrClient struct {
	Client     *solr.JSONClient
	Collection string
}

func NewSolrClient(host string, port int, collection string) *SolrClient {
	logger.Debug(fmt.Sprintf("%s:%d", host, port))
	Client := solr.NewJSONClient(fmt.Sprintf("http://%s:%d", host, port))
	return &SolrClient{
		Client:     Client,
		Collection: collection,
	}
}
func (sc *SolrClient) GetQuery(query string) (dtos.ItemsDTO, e.ApiError) {
	var response dtos.SolrResponseDto
	var itemsDto dtos.ItemsDTO
	url := "http://localhost:8983/solr/publicaciones/select?" + query + "=&defType=lucene&indent=true&q.op=OR&q=*%3A*"
	fmt.Println(url)
	q, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return itemsDto, e.NewInternalServerApiError("error hacendo query a solr", err)
	}

	var body []byte
	q.Body.Read(body)
	err = json.Unmarshal(body, &response)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return itemsDto, e.NewInternalServerApiError("error in unmarshal", err)

	}

	itemsDto = response.Response.Docs
	log.Println(err)
	return itemsDto, nil

}

func (sc *SolrClient) Update() error {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"COLA2", false, false, false, false, nil,
	)

	failOnError(err, "Failed to declare a queue")
	msgs, error2 := ch.Consume(
		q.Name, "", true, false, false, true, nil)
	failOnError(error2, "Failed to register consumer")

	var error error
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received message %s", d.Body)
			msg := string(d.Body)
			//solucionar
			url := "http://localhost:8090/items/" + msg
			fmt.Println(url)
			resp, err := http.Get(url)

			if err != nil {
				log.Fatalln(err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			//Convert the body to type string
			sb := string(body)
			log.Printf(sb)
			json.Marshal(sb)

			json_bytes := []byte(sb)
			var info Publi
			json.Unmarshal(json_bytes, &info)

			//Post en Solr
			collection := "publicaciones"
			//baseURL := "http://localhost:8983"
			//client := solr.NewJSONClient(baseURL)
			docs := []solr.M{
				{"id": msg, "title": info.Title, "description": info.Description, "city": info.City, "state": info.State, "image": info.Image},
			}
			buf := &bytes.Buffer{}
			error = json.NewEncoder(buf).Encode(docs)

			ctx := context.Background()

			_, error = sc.Client.Update(ctx, collection, solr.JSON, buf)

			error = sc.Client.Commit(ctx, collection)

		}
	}()
	<-forever

	return error
}
