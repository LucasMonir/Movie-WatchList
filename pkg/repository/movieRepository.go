package repository

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/utils"
	"os"
)

// Temporary file storage... migrate to mysql
func CheckStorageCreated() bool {
	jsonPath := utils.GetTaskFilePath()

	if !CheckStorageExists(jsonPath) {
		initStorage(jsonPath)
	}

	return true
}

func CheckStorageExists(jsonPath string) bool {
	info, err := os.Stat(jsonPath)
	return !(os.IsNotExist(err) || utils.CheckError(err) || info.IsDir() || utils.IsFileEmpty(jsonPath))
}

func initStorage(jsonPath string) {

	err := os.WriteFile(jsonPath, []byte("[]"), fs.ModePerm)
	if utils.CheckError(err) {
		fmt.Println(err.Error())
	}
}

func ReadMovies() []models.Movie {
	var movies []models.Movie
	jsonPath := utils.GetTaskFilePath()

	CheckStorageCreated()

	jsonFile, err := os.ReadFile(jsonPath)

	if utils.CheckError(err) {
		return movies
	}

	err = json.Unmarshal(jsonFile, &movies)

	if utils.CheckError(err) {
		return movies
	}

	return movies
}

func CreateMovie(name string) bool {
	movie, err := models.NewMovie(name)

	if utils.CheckError(err) {
		return false
	}

	movies := ReadMovies()

	if len(movies) != 0 {
		movie.Id = getLastId(movies)
	} else {
		movie.Id = 1
	}

	movies = append(movies, movie)
	json, err := json.MarshalIndent(movies, "", "	")

	if utils.CheckError(err) {
		fmt.Println("Error Marshaling the JSON...")
		return false
	}

	err = writeToStorage(json)

	if utils.CheckError(err) {
		fmt.Println("Error while writing movie to storage")
		return false
	}

	fmt.Println("Movie added successfully!")
	return true
}

func getLastId(movies []models.Movie) int {
	return movies[len(movies)-1].Id + 1
}

func writeToStorage(movies []byte) error {
	return os.WriteFile(utils.GetTaskFilePath(), movies, 0644)
}
