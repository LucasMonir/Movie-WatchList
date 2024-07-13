package queue

import (
	"context"
	"movie-watchlist/pkg/utils"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SendLogToServer(message string) bool {
	connection, err := amqp.Dial(os.Getenv("RABBIT_MQ_PROD"))

	if utils.CheckError(err) {
		panic("Error while starting queue service")
	}

	channel, err := connection.Channel()

	if utils.CheckError(err) {
		panic("Error while starting queue service")
	}

	defer channel.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = channel.PublishWithContext(ctx, "movie-watchlist-log", "log", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})

	if utils.CheckError(err) {
		panic("Error while starting queue service")
	}

	return true
}
