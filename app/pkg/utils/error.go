package utils

import "fmt"

func CheckError(err error) bool {
	return err != nil
}

func CheckAndPrintError(err error) bool {
	if !CheckError(err) {
		return false
	}

	fmt.Println(err.Error())
	return true
}
