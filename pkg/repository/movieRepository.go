package repository

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/utils"
	"os"
	"slices"
)

// Temporary file storage... migrate to mysql
func CheckStorageCreated() bool {
	jsonPath := utils.GetTaskFilePath()

	if !CheckStorageExists(jsonPath) {
		initStorage(jsonPath)
	}

	return true
}

// Temporary file storage... migrate to mysql
func CheckStorageExists(jsonPath string) bool {
	info, err := os.Stat(jsonPath)
	return !(os.IsNotExist(err) || utils.CheckError(err) || info.IsDir() || utils.IsFileEmpty(jsonPath))
}

// Temporary file storage... migrate to mysql
func initStorage(jsonPath string) {

	err := os.WriteFile(jsonPath, []byte("[]"), fs.ModePerm)
	if utils.CheckError(err) {
		fmt.Println(err.Error())
	}
}

func ReadMovies() ([]models.Movie, error) {
	var movies []models.Movie
	jsonPath := utils.GetTaskFilePath()

	CheckStorageCreated()

	jsonFile, err := os.ReadFile(jsonPath)

	if utils.CheckError(err) {
		return movies, err
	}

	err = json.Unmarshal(jsonFile, &movies)

	if utils.CheckError(err) {
		return movies, err
	}

	return movies, nil
}

func CreateMovie(name string) bool {
	movie, err := models.NewMovie(name)

	if utils.CheckError(err) {
		return false
	}

	movies, err := ReadMovies()

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

func RateMovie(id int, rating float32) bool {
	fmt.Println(id, rating)
	movies, err := ReadMovies()

	if utils.CheckError(err) {
		fmt.Println("Error while reading movies")
		return false
	}

	movieIndex := slices.IndexFunc[[]models.Movie](movies, func(m models.Movie) bool { return m.Id == id })
	movie := movies[movieIndex]

	fmt.Println(movie)

	movie.Rating = rating
	movies[movieIndex] = movie

	json, err := json.MarshalIndent(movies, "", "	")

	if utils.CheckError(err) {
		fmt.Println("Error serializing to json")
		return false
	}

	err = writeToStorage(json)

	if utils.CheckError(err) {
		fmt.Println("Error serializing to json")
		return false
	}

	return true
}

func getLastId(movies []models.Movie) int {
	return movies[len(movies)-1].Id + 1
}

func writeToStorage(movies []byte) error {
	return os.WriteFile(utils.GetTaskFilePath(), movies, 0644)
}
