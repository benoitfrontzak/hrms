package consumer

import (
	"bytes"
	"log"
	"net/http"
)

// Listen watch the consumer's queue and consume events
func (c *Consumer) Listen() error {
	// opening a channel over the connection established to interact with RabbitMQ
	ch, err := c.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// tell the server to deliver us the messages from the queue.
	// Since it will push us messages asynchronously,
	// we will read the messages from a channel (returned by amqp::Consume) in a goroutine.
	msgs, err := ch.Consume(
		c.queueName, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		log.Println("failed to register a consumer: ", err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message from: %s", d.RoutingKey)
			err := sendEmail(d.Body)
			if err != nil {
				log.Println("error while sending emnail: ", err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}

// sendEmail send the payload email to mailer-service
func sendEmail(entry []byte) error {

	mailerServiceURL := "http://mailer-service:8888/sendEmail"

	response, err := http.Post(mailerServiceURL, "application/json", bytes.NewBuffer(entry))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// request, err := http.NewRequest("POST", mailerServiceURL, bytes.NewBuffer(entry))
	// if err != nil {
	// 	return err
	// }

	// request.Header.Set("Content-Type", "application/json")
	// client := &http.Client{}
	// response, err := client.Do(request)
	// if err != nil {
	// 	return err
	// }
	// defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return err
	}

	log.Println("http response from mailer-service: ", response.StatusCode)

	return nil
}
