package queue

import (
	"context"
	"fmt"
	"movie-watchlist/pkg/utils"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func CheckRmqStarted() {
	var err error = fmt.Errorf("Placeholder")
	var connection *amqp.Connection

	for err != nil {
		time.Sleep(10 * time.Second)
		connection, err = amqp.Dial(os.Getenv("RABBIT_MQ_PROD"))
	}

	defer connection.Close()
	fmt.Println("RabbitMQ connection started successfully!")
}

func SendLogToServer(message string) bool {
	var connection *amqp.Connection

	connection, err := amqp.Dial(os.Getenv("RABBIT_MQ_PROD"))

	if utils.CheckAndPrintError(err) {
		return false
	}

	defer connection.Close()

	channel, err := connection.Channel()

	if utils.CheckAndPrintError(err) {
		return false
	}

	err = channel.ExchangeDeclare(
		"logs",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	if utils.CheckAndPrintError(err) {
		return false
	}

	defer channel.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err = channel.PublishWithContext(
		ctx,
		"logs",
		"logs",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	return !utils.CheckAndPrintError(err)
}

func CheckAndLogError(err error) bool {
	if !utils.CheckError(err) {
		return false
	}

	SendLogToServer(err.Error())
	return true
}
