package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
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

func ConsumerSolr(msg string) error {

	url := "https://localhost:8090/items/" + msg
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
	baseURL := "http://localhost:8983"
	client := solr.NewJSONClient(baseURL)
	docs := []solr.M{
		{"id": msg, "title": info.Title, "description": info.Description, "city": info.City, "state": info.State, "image": info.Image},
	}
	buf := &bytes.Buffer{}
	error := json.NewEncoder(buf).Encode(docs)

	ctx := context.Background()

	_, error = client.Update(ctx, collection, solr.JSON, buf)

	error = client.Commit(ctx, collection)

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
	_, error2 := ch.Consume(
		q.Name, "", true, false, false, true, nil)
	failOnError(error2, "Failed to register consumer")

	return error
}
