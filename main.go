package main

import (
	"fmt"
	"log"
	mailer "logimailservice/mail"
	"logimailservice/util"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
	}
}

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("config file could not be loaded: ", err)
	}

	conn, err := amqp.Dial(config.Amqp_address)
	failOnError(err, "failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"mailch",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to reegister the consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			fmt.Printf("Received message: %s", d.Body)
			err := mailer.Smail()
			failOnError(err, "Failed to send message")
		}
	}()

	log.Printf(" [*] Waiting for messages. Press ctrl+c to exit")
	<-forever

}
