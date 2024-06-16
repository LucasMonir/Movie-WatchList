package controller

import (
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/repository"
	"movie-watchlist/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovies(context *gin.Context) {
	movies := repository.ReadMovies()
	context.IndentedJSON(http.StatusOK, movies)
}

func GetMovie(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, nil)
}

func AddMovie(context *gin.Context) {
	var movie = models.Movie{}

	err := context.BindJSON(&movie)

	if utils.CheckError(err) {
		context.IndentedJSON(http.StatusBadRequest, false)
	}

	result := repository.CreateMovie(movie.Name)
	context.IndentedJSON(http.StatusOK, result)
}

func DeleteMovie(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, nil)
}

func RateMovie(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, nil)
}
