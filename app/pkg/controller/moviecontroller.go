package controller

import (
	"fmt"
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/queue"
	"movie-watchlist/pkg/repository"
	"movie-watchlist/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMovies(context *gin.Context) {
	queue.SendLogToServer("[GET] GetMovies")

	movies, err := repository.ReadMovies()

	if utils.CheckError(err) {
		fmt.Println("Error while reading movies")
		context.IndentedJSON(http.StatusInternalServerError, false)
		return
	}

	context.IndentedJSON(http.StatusOK, movies)
}

func GetMovie(context *gin.Context) {
	queue.SendLogToServer("[GET] GetMovie")
	id, err := getIdFromParams(context.Params)

	if utils.CheckError(err) {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	movie, err := repository.ReadMovie(id)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, fmt.Errorf("no movie found with the id: %d", id).Error())
		return
	}

	context.IndentedJSON(http.StatusOK, movie)
}

func AddMovie(context *gin.Context) {
	queue.SendLogToServer("[POST] AddMovie")

	var movie = models.Movie{}

	err := context.BindJSON(&movie)

	if utils.CheckError(err) {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	result, id, err := repository.CreateMovie(movie.Name)

	if utils.CheckError(err) || !result {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(http.StatusOK, id)
}

func RateMovie(context *gin.Context) {
	queue.SendLogToServer("[PATCH] RateMovie")

	id, err := getIdFromParams(context.Params)

	if utils.CheckError(err) {
		context.IndentedJSON(http.StatusBadRequest, false)
		return
	}

	rating, err := strconv.ParseFloat(context.Query("rating"), 32)

	if utils.CheckError(err) {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := repository.RateMovie(id, float32(rating))

	if !result || err != nil {
		fmt.Println(err.Error())
		context.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.IndentedJSON(http.StatusOK, true)
}

func DeleteMovie(context *gin.Context) {
	queue.SendLogToServer("[DELETE] DeleteMovie")

	id, err := getIdFromParams(context.Params)

	if utils.CheckError(err) {
		context.IndentedJSON(http.StatusBadRequest, "Invalid ID parameter")
		return
	}

	result, err := repository.DeleteMovie(id)

	if !result || err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(http.StatusOK, result)
}

func getIdFromParams(params gin.Params) (int, error) {
	idParam := params.ByName("id")

	return strconv.Atoi(idParam)
}
