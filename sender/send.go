package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s,%s", msg, err)
	}
}

func main() {

	conn,err :=  amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err,"failed to connect to rabbitmq")
	defer conn.Close()

	ch,err :=conn.Channel()
	failOnError(err,"failed to open a channel")
	defer ch.Close()

	q,err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
		)

	failOnError(err,"failed to declare a queue")

	body := "hello world"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			Headers:         nil,
			ContentType:     "",
			ContentEncoding: "",
			DeliveryMode:    0,
			Priority:        0,
			CorrelationId:   "",
			ReplyTo:         "",
			Expiration:      "",
			MessageId:       "",
			Timestamp:       time.Time{},
			Type:            "",
			UserId:          "",
			AppId:           "",
			Body:            []byte(body),
		})

	log.Printf("[x] sent %s",body)
	failOnError(err , "failed to publish a message")


}
