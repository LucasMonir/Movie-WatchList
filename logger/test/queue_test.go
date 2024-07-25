package test

import (
	"fmt"
	"logger/pkg/queue"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectToRabbitMQMustReturnConnection(t *testing.T) {
	connection, err := queue.ConnectRabbitMQ()

	assert.False(t, err == nil, fmt.Sprint("No error was expected, found: %v", err.Error()))

	assert.False(t, connection.IsClosed(), "Connection must be open, found to be closed")
}
