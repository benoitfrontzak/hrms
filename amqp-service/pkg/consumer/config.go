package consumer

import amqp "github.com/rabbitmq/amqp091-go"

// Consumer is the type which holds the amqp connection and queue name
type Consumer struct {
	connection *amqp.Connection
	queueName  string
}

// NewConsumer returns a Consumer type
func NewConsumer(conn *amqp.Connection, queue string) *Consumer {
	return &Consumer{
		connection: conn,
		queueName:  queue,
	}
}
