package controller

import (
	"movie-watchlist/pkg/repository"
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

func InsertMovie(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, nil)

}

func DeleteMovie(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, nil)
}

func RateMovie(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, nil)
}
