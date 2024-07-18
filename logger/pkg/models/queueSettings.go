package models

import "github.com/rabbitmq/amqp091-go"

type RmqConfig struct {
	kind       bool
	internal   bool
	key        string
	durable    bool
	exclusive  bool
	noWait     bool
	autoDelete bool
	queueName  string
	exchange   string
	template   amqp091.Table
}
