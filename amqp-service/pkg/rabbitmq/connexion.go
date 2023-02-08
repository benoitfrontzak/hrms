package rabbitmq

import (
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Connect dials up rabbitmq server on port 5672,
// it returns a new Connection over TCP
// (HTTP management UI on port 15672)
func Connect() (*amqp.Connection, error) {
	var counts int64 // counts holds the number of connexion attempt to the rabbitmq server
	var connection *amqp.Connection

	// AMQP URI format
	dsn := os.Getenv("DSN")

	for {
		// Dial accepts a string in the AMQP URI format and returns a new Connection over TCP using PlainAuth
		// DSN is set as environment variable in docker-compose.yml
		c, err := amqp.Dial(dsn)
		if err != nil {
			fmt.Println("RabbitMQ not yet ready, backing off for two seconds...")
			counts++
		} else {
			log.Println("Connected to RabbitMQ!")
			connection = c
			break
		}
		// exit program after 10 attemps
		if counts > 10 {
			fmt.Println(err)
			return nil, err
		}
		// Wait for 2 seconds before dialing again
		time.Sleep(2 * time.Second)
		continue
	}

	return connection, nil
}
