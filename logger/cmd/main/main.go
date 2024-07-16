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

	var err error = fmt.Errorf("Placeholder")
	var connection *amqp.Connection

	for err != nil {
		connection, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")

		if utils.CheckError(err) {
			fmt.Println("Error while connecting to rabbitmq")
			fmt.Print(err.Error())
		}
	}

	defer connection.Close()

	channel, err := connection.Channel()

	if utils.CheckError(err) {
		fmt.Println("Error while starting channel service")
		fmt.Print(err.Error())
	}

	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"logs",
		true,
		false,
		false,
		false,
		nil,
	)

	if utils.CheckError(err) {
		fmt.Println("Error while declaring queue")
		fmt.Print(err.Error())
	}

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
		fmt.Print(err.Error())
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
