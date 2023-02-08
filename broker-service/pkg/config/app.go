package config

import amqp "github.com/rabbitmq/amqp091-go"

// AppConfig holds the application config
type AppConfig struct {
	Rabbit *amqp.Connection
}
