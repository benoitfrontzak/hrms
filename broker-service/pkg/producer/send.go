package producer

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Producer is the type which holds the amqp connection
type Producer struct {
	connection *amqp.Connection
}

// NewProducer returns a Producer type
func NewProducer(conn *amqp.Connection) *Producer {
	return &Producer{
		connection: conn,
	}
}

// PushMail publishes the queue with the message
func (p *Producer) PushMail(queue string, msg []byte) error {
	// opening a channel over the connection established to interact with RabbitMQ
	channel, err := p.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	// Create a context for publishing
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// publishing the message
	err = channel.PublishWithContext(ctx,
		"",    // exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
