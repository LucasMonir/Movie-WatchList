package utils

import (
	"encoding/json"
	models "movie-watchlist/pkg/models"
	"os"
)

func GetTaskFilePath() string {
	configsJson, err := os.ReadFile("../../configs.json")

	if CheckError(err) {
		return ""
	}

	var path models.Path

	err = json.Unmarshal(configsJson, &path)

	if CheckError(err) {
		return ""
	}

	return "../../" + path.StoragePath
}

func IsFileEmpty(jsonPath string) bool {
	file, err := os.ReadFile(jsonPath)

	if CheckError(err) {
		return true
	}

	if len(file) == 0 {
		return true
	}

	return false
}
