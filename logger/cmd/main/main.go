package main

import (
	"fmt"
	"logger/pkg/queue"
	"os"
)

func main() {
	checkEnvironment()
	queue.ConsumeMessages()
}

func checkEnvironment() {
	fmt.Println(os.Getenv("RABBIT_MQ_PROD"))
}
