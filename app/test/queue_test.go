package test

import (
	"movie-watchlist/pkg/queue"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendLogToServerMustSendMessage(t *testing.T) {
	result := queue.SendLogToServer("Test")
	assert.True(t, result, "Result must return true")
}
