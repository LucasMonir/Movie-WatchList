package utils

import "fmt"

func CheckError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}

	return false
}
