package main

import (
	"fmt"
	"log"
	"logger/pkg/utils"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	checkEnvironment()
	connection, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")

	if utils.CheckError(err) {
		fmt.Println("Error while starting queue service")
	}

	defer connection.Close()

	channel, err := connection.Channel()

	if utils.CheckError(err) {
		fmt.Println("Error while starting channel service")
	}

	defer channel.Close()

	err = channel.ExchangeDeclare(
		"movie-watchlist-log",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	fmt.Println("declared exhc")

	if utils.CheckError(err) {
		fmt.Println("Error while declaring exchange")
	}

	queue, err := channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)

	fmt.Println("declared wqueue")

	if utils.CheckError(err) {
		fmt.Println("Error while declaring queue")
	}

	channel.QueueBind(queue.Name, "log", "movie-watchlist-log", false, nil)

	fmt.Println("bound queue")

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
		fmt.Println("Error while consuming messages")
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

func checkEnvironment() {
	fmt.Println(os.Getenv("RABBIT_MQ_PROD"))
}
