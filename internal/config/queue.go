package config

import (
	"crypto/tls"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	Conn         *amqp.Connection
	ExchangeName string
	QueueName    string
	RoutingKey   string
}

func SetupRabbitMQConnection(config RabbitMQConfig, serverConfig ServerConfig) (*Queue, error) {
	var conn *amqp.Connection
	var err error

	if serverConfig.Env == "local" {
		connectionUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Host, config.Port)
		conn, err = amqp.Dial(connectionUrl)
	} else {
		connectionUrl := fmt.Sprintf("amqps://%s:%s@%s:%d/", config.Username, config.Password, config.Host, config.Port)

		tlsConfig := &tls.Config{
			MinVersion: tls.VersionTLS12,
		}

		conn, err = amqp.DialTLS(connectionUrl, tlsConfig)
	}
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
		QueueName:    config.QueueName,
		RoutingKey:   config.RoutingKey,
	}, nil
}
