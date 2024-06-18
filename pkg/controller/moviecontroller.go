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
		fmt.Println("Error while converting ID parameter to int")
		context.IndentedJSON(http.StatusBadRequest, false)
		return
	}

	movies, err := repository.ReadMovies()

	if utils.CheckError(err) {
		fmt.Println("Error while reading movies")
		context.IndentedJSON(http.StatusInternalServerError, false)
		return
	}

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
	id, err := getIdFromParams(context.Params)

	if utils.CheckError(err) {
		fmt.Println("Invalid ID parameter")
		context.IndentedJSON(http.StatusBadRequest, false)
		return
	}

	rating, err := strconv.ParseFloat(context.Query("rating"), 32)

	if utils.CheckError(err) {
		fmt.Println("Invalid Rating parameter")
		context.IndentedJSON(http.StatusBadRequest, false)
		return
	}

	result := repository.RateMovie(id, float32(rating))

	if !result {
		fmt.Println("Error updating movie rating...")
		context.IndentedJSON(http.StatusInternalServerError, false)
		return
	}

	context.IndentedJSON(http.StatusOK, "Rating added sucessfully!")
}

func getIdFromParams(params gin.Params) (int, error) {
	idParam := params.ByName("id")

	return strconv.Atoi(idParam)
}
