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
		connection, err = amqp.Dial(os.Getenv("RABBIT_MQ_PROD"))

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
		fmt.Println("Error while declaring exchange")
		fmt.Print(err.Error())
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
		fmt.Println("Error while declaring queue")
		fmt.Print(err.Error())
	}

	err = channel.QueueBind(
		queue.Name,
		"logs",
		"logs",
		false,
		nil,
	)

	if utils.CheckError(err) {
		fmt.Println("Error while binding queue")
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
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func checkEnvironment() {
	fmt.Println(os.Getenv("RABBIT_MQ_PROD"))
}
