package main

import (
	"amqp/pkg/consumer"
	"amqp/pkg/rabbitmq"
	"log"
	"os"
)

func main() {
	// connect to rabbitmq server
	rabbitConn, err := rabbitmq.Connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	// create mailer consumer
	queue := "mailer"
	mailer := consumer.NewConsumer(rabbitConn, queue)

	// declare mailer queue
	err = mailer.Declare()
	if err != nil {
		log.Println(err)
	}

	// start listening for messages
	log.Println("Listening for and consuming RabbitMQ messages...")

	// watch mailer queue and consume events
	err = mailer.Listen()
	if err != nil {
		log.Println(err)
	}

}
