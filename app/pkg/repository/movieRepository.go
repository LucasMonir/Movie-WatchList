package repository

import (
	"fmt"
	"movie-watchlist/pkg/db"
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/utils"
)

// // Temporary file storage... migrate to mysql
// func CheckStorageCreated() bool {
// 	jsonPath := utils.GetTaskFilePath()

// 	if !CheckStorageExists(jsonPath) {
// 		initStorage(jsonPath)
// 	}

// 	return true
// }

// // Temporary file storage... migrate to mysql
// func CheckStorageExists(jsonPath string) bool {
// 	info, err := os.Stat(jsonPath)
// 	return !(os.IsNotExist(err) || utils.CheckError(err) || info.IsDir() || utils.IsFileEmpty(jsonPath))
// }

// // Temporary file storage... migrate to mysql
// func initStorage(jsonPath string) {

// 	err := os.WriteFile(jsonPath, []byte("[]"), fs.ModePerm)
// 	if utils.CheckError(err) {
// 		fmt.Println(err.Error())
// 	}
// }

// func CreateMovie(name string) (bool, error) {
// 	movie, err := models.NewMovie(name)

// 	if utils.CheckError(err) {
// 		return false, fmt.Errorf("error creating new movie, check the parameters")
// 	}

// 	movies, err := ReadMovies()

// 	if utils.CheckError(err) {
// 		return false, fmt.Errorf("error while reading stored movies")
// 	}

// 	if len(movies) != 0 {
// 		movie.Id = getLastId(movies)
// 	} else {
// 		movie.Id = 1
// 	}

// 	movies = append(movies, movie)
// 	json, err := json.MarshalIndent(movies, "", "	")

// 	if utils.CheckError(err) {
// 		return false, fmt.Errorf("error Marshaling the JSON")
// 	}

// 	err = writeToStorage(json)

// 	if utils.CheckError(err) {
// 		return false, fmt.Errorf("error writing movie to storage")
// 	}

// 	fmt.Println("Movie added successfully!")
// 	return true, nil
// }

// func RateMovie(id int, rating float32) (bool, error) {
// 	if !checkRating(rating) {
// 		return false, fmt.Errorf("invalid Rating: %f", rating)
// 	}

// 	movies, err := ReadMovies()

// 	if utils.CheckError(err) {
// 		return false, fmt.Errorf("error while reading movies")
// 	}

// 	movieIndex := slices.IndexFunc(movies, func(m models.Movie) bool { return m.Id == id })
// 	movie := movies[movieIndex]

// 	fmt.Println(movie)

// 	movie.Rating = rating
// 	movies[movieIndex] = movie

// 	json, err := json.MarshalIndent(movies, "", "	")

// 	if utils.CheckError(err) {
// 		return false, fmt.Errorf("error serializing to json")
// 	}

// 	err = writeToStorage(json)

// 	if utils.CheckError(err) {
// 		return false, fmt.Errorf("error serializing to json")
// 	}

// 	return true, nil
// }

// func getLastId(movies []models.Movie) int {
// 	return movies[len(movies)-1].Id + 1
// }

// func writeToStorage(movies []byte) error {
// 	return os.WriteFile(utils.GetTaskFilePath(), movies, 0644)
// }

// func checkRating(rating float32) bool {
// 	return rating >= 0 && rating <= 10
// }

func ReadMovies() ([]models.Movie, error) {
	var movies []models.Movie

	db := db.GetConnection()

	defer db.Close()

	err := db.Model(&movies).Select()

	if utils.CheckError(err) {
		return nil, err
	}

	return movies, nil
}

func CreateMovie(name string) (bool, error) {
	movie, err := models.NewMovie(name)

	if utils.CheckError(err) {
		return false, fmt.Errorf("error creating new movie, check the parameters")
	}

	db := db.GetConnection()

	result, err := db.Model(&movie).Insert()

	if utils.CheckError(err) || result.RowsAffected() == 0 {
		return false, fmt.Errorf("error inserting movie into database")
	}

	return true, nil
}
