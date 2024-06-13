package repository

import (
	"fmt"
	"io/fs"
	"movie-watchlist/pkg/utils"
	"os"
)

// Temporary file storage... migrate to mysql
func CheckStorageCreated() bool {
	jsonPath := utils.GetTaskFilePath()

	info, err := os.Stat(jsonPath)

	if os.IsNotExist(err) || utils.CheckError(err) || info.IsDir() || utils.IsFileEmpty(jsonPath) {
		fmt.Println("Creating task file")
		initStorage(jsonPath)
	}

	return true
}

func initStorage(jsonPath string) {

	err := os.WriteFile(jsonPath, []byte("[]"), fs.ModePerm)
	if utils.CheckError(err) {
		fmt.Println(err.Error())
	}
}
