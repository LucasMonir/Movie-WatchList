package repository

import (
	"fmt"
	"movie-watchlist/pkg/db"
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/utils"
)

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

func ReadMovie(id int) (models.Movie, error) {
	var movie models.Movie

	db := db.GetConnection()

	defer db.Close()

	err := db.Model(&movie).Where("id = ?", id).Select()

	if utils.CheckError(err) {
		return movie, err
	}

	return movie, nil
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

func RateMovie(id int, rating float32) (bool, error) {
	// if !checkRating(rating) {
	// 	return false, fmt.Errorf("invalid Rating: %f", rating)
	// }

	// movies, err := ReadMovie()

	// if utils.CheckError(err) {
	// 	return false, fmt.Errorf("error while reading movies")
	// }

	// movieIndex := slices.IndexFunc(movies, func(m models.Movie) bool { return m.Id == id })
	// movie := movies[movieIndex]

	// fmt.Println(movie)

	// movie.Rating = rating
	// movies[movieIndex] = movie

	// json, err := json.MarshalIndent(movies, "", "	")

	// if utils.CheckError(err) {
	// 	return false, fmt.Errorf("error serializing to json")
	// }

	// err = writeToStorage(json)

	// if utils.CheckError(err) {
	// 	return false, fmt.Errorf("error serializing to json")
	// }

	return true, nil
}

func checkRating(rating float32) bool {
	return rating >= 0 && rating <= 10
}
