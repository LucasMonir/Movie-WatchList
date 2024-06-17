package controller

import (
	"fmt"
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/repository"
	"movie-watchlist/pkg/utils"
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMovies(context *gin.Context) {
	movies := repository.ReadMovies()
	context.IndentedJSON(http.StatusOK, movies)
}

func GetMovie(context *gin.Context) {
	idParam := context.Param("id")

	id, err := strconv.Atoi(idParam)

	if utils.CheckError(err) {
		fmt.Println("Error while converting ID parameter to int")
		context.IndentedJSON(http.StatusBadRequest, false)
		return
	}

	movies := repository.ReadMovies()

	if utils.CheckError(err) {
		fmt.Println("Error while reading movies")
		context.IndentedJSON(http.StatusBadRequest, false)
		return
	}

	fmt.Println(id)
	movieIndex := slices.IndexFunc[[]models.Movie](movies, func(m models.Movie) bool { return m.Id == id })

	fmt.Println(movieIndex)
	context.IndentedJSON(http.StatusOK, movies[movieIndex])
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
