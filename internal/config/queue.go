package config

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Queue struct {
	Conn         *amqp.Connection
	ExchangeName string
	queueName    string
	RoutingKey   string
}

func SetupRabbitMQConnection(config RabbitMQConfig) (*Queue, error) {
	connectionUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Host, config.Port)

	conn, err := amqp.Dial(connectionUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		config.ExchangeName,
		config.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	_, err = ch.QueueDeclare(
		config.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	err = ch.QueueBind(
		config.QueueName,
		config.RoutingKey,
		config.ExchangeName,
		false,
		nil,
	)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	err = ch.Qos(
		config.PrefetchCount,
		0,
		false,
	)

	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to set QoS: %w", err)
	}

	log.Printf("RabbitMQ connection established and queue configured successfully")
	return &Queue{Conn: conn,
		ExchangeName: config.ExchangeName,
		queueName:    config.QueueName,
		RoutingKey:   config.RoutingKey,
	}, nil
}
