package utils

func CheckError(err error) bool {
	if err != nil {
		LogError(err)
		return true
	}

	return false
}
