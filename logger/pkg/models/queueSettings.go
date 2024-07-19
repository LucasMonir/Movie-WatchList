package models

import "github.com/rabbitmq/amqp091-go"

type RmqConfig struct {
	Kind       string
	Internal   bool
	AutoAck    bool
	Key        string
	Durable    bool
	Exclusive  bool
	NoWait     bool
	AutoDelete bool
	QueueName  string
	Exchange   string
	Consumer   string
	NoLocal    bool
	Template   amqp091.Table
}

func GetRmqConfig() RmqConfig {
	return RmqConfig{
		Kind:       "direct",
		Internal:   false,
		AutoAck:    true,
		Key:        "logs",
		Durable:    false,
		NoWait:     false,
		Exchange:   "logs",
		AutoDelete: false,
		Template:   nil,
		Exclusive:  true,
		QueueName:  "",
		Consumer:   "",
		NoLocal:    false,
	}
}
