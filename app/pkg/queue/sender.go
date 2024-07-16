package queue

import (
	"context"
	"fmt"
	"movie-watchlist/pkg/utils"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SendLogToServer(message string) bool {
	var err error = fmt.Errorf("Placeholder")
	var connection *amqp.Connection

	for err != nil {
		connection, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")

		if utils.CheckError(err) {
			fmt.Println("Error while connecting to rabbitmq")
			fmt.Print(err.Error())
			return false
		}
	}

	defer connection.Close()

	channel, err := connection.Channel()

	if utils.CheckError(err) {
		fmt.Println("Error while opening channel")
		fmt.Print(err.Error())
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

	if utils.CheckError(err) {
		fmt.Println("Error while creating queue")
		fmt.Print(err.Error())
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

	if utils.CheckError(err) {
		fmt.Println("Error while publishing message")
		fmt.Print(err.Error())
		return false
	}

	return true
}
