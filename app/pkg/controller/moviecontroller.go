package controller

import (
	"fmt"
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/repository"
	"movie-watchlist/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMovies(context *gin.Context) {
	movies, err := repository.ReadMovies()

	if utils.CheckError(err) {
		fmt.Println("Error while reading movies")
		context.IndentedJSON(http.StatusInternalServerError, false)
		return
	}

	context.IndentedJSON(http.StatusOK, movies)
}

func GetMovie(context *gin.Context) {
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
	var movie = models.Movie{}

	err := context.BindJSON(&movie)

	if utils.CheckError(err) {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := repository.CreateMovie(movie.Name)

	if utils.CheckError(err) {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(http.StatusOK, result)
}

func RateMovie(context *gin.Context) {
	id, err := getIdFromParams(context.Params)

	if utils.CheckError(err) {
		fmt.Println("Invalid ID parameter")
		context.IndentedJSON(http.StatusBadRequest, false)
		return
	}

	rating, err := strconv.ParseFloat(context.Query("rating"), 32)

	if utils.CheckError(err) {
		fmt.Println("Invalid Rating parameter")
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := repository.RateMovie(id, float32(rating))

	if !result || err != nil {
		fmt.Println(err.Error())
		context.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.IndentedJSON(http.StatusOK, "Rating added sucessfully!")
}

func getIdFromParams(params gin.Params) (int, error) {
	idParam := params.ByName("id")

	return strconv.Atoi(idParam)
}

func DeleteMovie(context *gin.Context) {
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
