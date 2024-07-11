package main

import (
	"fmt"
	"log"
	"logger/pkg/utils"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if utils.CheckError(err) {
		panic("Error while starting queue service")
	}

	channel, err := connection.Channel()

	if utils.CheckError(err) {
		panic("Error while starting queue service")
	}

	err = channel.ExchangeDeclare(
		"movie-watchlist-log",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	if utils.CheckError(err) {
		panic("Error while starting queue service")
	}

	queue, err := channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)

	if utils.CheckError(err) {
		panic("Error while starting queue service")
	}

	channel.QueueBind(queue.Name, "log", "movie-watchlist-log", false, nil)

	msgs, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if utils.CheckError(err) {
		panic("Error while starting queue service")
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			fmt.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
