package consumer

// Declare declares a queue for us to consume to
func (c *Consumer) Declare() error {
	// opening a channel over the connection established to interact with RabbitMQ
	ch, err := c.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// declaring queue with its properties over the channel opened
	_, err = ch.QueueDeclare(
		c.queueName, // name?
		false,       // durable?
		false,       // delete when unused?
		true,        // exclusive?
		false,       // no-wait?
		nil,         // arguments?
	)
	if err != nil {
		return err
	}

	return nil
}
