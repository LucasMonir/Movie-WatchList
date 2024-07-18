package queue

import (
	"fmt"
	"log"
	"logger/pkg/utils"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeMessages() {
	connection, err := ConnectRabbitMQ()

	if utils.CheckError(err) {
		fmt.Println("Error connecting to rabbitmq")
		fmt.Print(err.Error())
	}

	defer connection.Close()

	channel, err := connection.Channel()

	if utils.CheckError(err) {
		fmt.Println("Error while starting channel service")
		fmt.Print(err.Error())
	}

	defer channel.Close()

	err = setUpExchange(channel)

	if utils.CheckError(err) {
		fmt.Println("Error while declaring exchange")
		fmt.Print(err.Error())
	}

	queue, err := setUpQueue(channel)

	if utils.CheckError(err) {
		fmt.Println("Error while declaring queue")
		fmt.Print(err.Error())
	}

	err = bindQueue(channel, queue.Name)

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

func ConnectRabbitMQ() (*amqp.Connection, error) {
	var connection *amqp.Connection
	var err error

	for attempts := 0; attempts < 10; attempts++ {
		connection, err = amqp.Dial(os.Getenv("RABBIT_MQ_PROD"))

		if err == nil {
			return connection, nil
		}

		log.Printf("Failed to connect to RabbitMQ, attempt %d: %v", attempts+1, err)
		time.Sleep(5 * time.Second)
	}

	return nil, fmt.Errorf("could not connect to RabbitMQ after multiple attempts: %v", err)
}

func setUpExchange(channel *amqp.Channel) error {
	err := channel.ExchangeDeclare(
		"logs",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	return err
}

func setUpQueue(channel *amqp.Channel) (amqp.Queue, error) {
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

		return amqp.Queue{}, err
	}

	return queue, err
}

func bindQueue(channel *amqp.Channel, queueName string) error {
	err := channel.QueueBind(
		queueName,
		"logs",
		"logs",
		false,
		nil,
	)

	return err
}
