package controller

import (
	"fmt"
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/queue"
	"movie-watchlist/pkg/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMovies(context *gin.Context) {
	movies, err := repository.ReadMovies()

	if queue.CheckAndLogError(err) {
		context.IndentedJSON(http.StatusInternalServerError, false)
		return
	}

	context.IndentedJSON(http.StatusOK, movies)
}

func GetMovie(context *gin.Context) {
	id, err := getIdFromParams(context.Params)

	if queue.CheckAndLogError(err) {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	movie, err := repository.ReadMovie(id)

	if queue.CheckAndLogError(err) {
		context.IndentedJSON(http.StatusInternalServerError, fmt.Errorf("no movie found with the id: %d", id).Error())
		return
	}

	context.IndentedJSON(http.StatusOK, movie)
}

func AddMovie(context *gin.Context) {
	var movie = models.Movie{}

	err := context.BindJSON(&movie)

	if queue.CheckAndLogError(err) {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	result, id, err := repository.CreateMovie(movie.Name)

	if queue.CheckAndLogError(err) || !result {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(http.StatusOK, id)
}

func RateMovie(context *gin.Context) {
	id, err := getIdFromParams(context.Params)

	if queue.CheckAndLogError(err) {
		context.IndentedJSON(http.StatusBadRequest, false)
		return
	}

	rating, err := strconv.ParseFloat(context.Query("rating"), 32)

	if queue.CheckAndLogError(err) {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := repository.RateMovie(id, float32(rating))

	if !result || err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.IndentedJSON(http.StatusOK, true)
}

func DeleteMovie(context *gin.Context) {
	id, err := getIdFromParams(context.Params)

	if queue.CheckAndLogError(err) {
		context.IndentedJSON(http.StatusBadRequest, "Invalid ID parameter")
		return
	}

	result, err := repository.DeleteMovie(id)

	if !result || queue.CheckAndLogError(err) {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(http.StatusOK, result)
}

func getIdFromParams(params gin.Params) (int, error) {
	idParam := params.ByName("id")

	return strconv.Atoi(idParam)
}
