package utils

import "fmt"

func CheckError(err error) bool {
	if err != nil {
		fmt.Print(err.Error())
		return true
	}

	return false
}
