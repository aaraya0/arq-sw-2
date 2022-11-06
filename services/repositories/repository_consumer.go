package repositories

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

func processIDs(msg []byte) {
	message := string(msg)
	url := "https://localhost:8090/items/" + message
	resp, err := http.Get(url)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	httpposturl := "https://localhost:8983/solr/publicaciones/update"
	var jsonData = body
	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

}

func Consume() {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"COLA", false, false, false, false, nil,
	)
	failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		q.Name, "", true, false, false, true, nil)
	failOnError(err, "Failed to register consumer")
	d := <-msgs
	go processIDs(d.Body)

}
