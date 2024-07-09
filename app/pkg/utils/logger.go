package utils

import (
	"fmt"
	"os"
	"strconv"
)

func LogError(err error) {
	if !checkLogingEnable() {
		return
	}

	fmt.Println(err.Error())
}

func LogInfo(info string) {
	if !checkLogingEnable() {
		return
	}

	fmt.Println(info)
}

func checkLogingEnable() bool {
	result := os.Getenv("LOG_ENABLED")
	value, err := strconv.ParseBool(result)

	if err != nil {
		fmt.Println("Error configuring logger")
		return false
	}

	return value
}
