package repository

import (
	"fmt"
	"movie-watchlist/pkg/db"
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/queue"
	"movie-watchlist/pkg/utils"
)

func ReadMovies() ([]models.Movie, error) {
	var movies []models.Movie

	db := db.GetConnection()
	defer db.Close()

	err := db.Model(&movies).Select()

	if queue.CheckAndLogError(err) {
		return movies, err
	}

	return movies, nil
}

func ReadMovie(id int) (models.Movie, error) {
	var movie models.Movie

	db := db.GetConnection()
	defer db.Close()

	err := db.Model(&movie).Where("id = ?", id).Select()

	if queue.CheckAndLogError(err) {
		return movie, err
	}

	return movie, nil
}

func CreateMovie(name string) (bool, int, error) {
	movie, err := models.NewMovie(name)

	if queue.CheckAndLogError(err) {
		return false, 0, fmt.Errorf("error creating new movie, check the parameters")
	}

	db := db.GetConnection()
	defer db.Close()

	result, err := db.Model(&movie).Insert()

	if checkRowsUpdated(err, result.RowsAffected()) {
		return false, 0, fmt.Errorf("error inserting movie into database")
	}

	return true, movie.Id, nil
}

func RateMovie(id int, rating float32) (bool, error) {
	if !checkRating(rating) {
		return false, fmt.Errorf("ratings must be equal or between 0 and 10: %f", rating)
	}

	movie := models.Movie{}
	db := db.GetConnection()
	defer db.Close()

	result, err := db.Model(&movie).Set("rating = ?", rating).Where("id = ?", id).Update()

	if checkRowsUpdated(err, result.RowsAffected()) {
		return false, fmt.Errorf("error while updating records: %s", err.Error())
	}

	return true, nil
}

func DeleteMovie(id int) (bool, error) {
	movie := models.Movie{}

	_, err := ReadMovie(id)

	if utils.CheckError(err) {
		return false, fmt.Errorf("movie not found")
	}

	db := db.GetConnection()
	defer db.Close()

	result, err := db.Model(&movie).Where("id = ?", id).Delete()

	if checkRowsUpdated(err, result.RowsAffected()) {
		return false, fmt.Errorf("error while deleting movie: %s", err.Error())
	}

	return true, nil
}

func checkRating(rating float32) bool {
	return rating >= 0 && rating <= 10
}

func checkRowsUpdated(err error, rows int) bool {
	return queue.CheckAndLogError(err) || rows == 0
}
