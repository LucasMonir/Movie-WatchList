package main

import (
	"fmt"
	"movie-watchlist/pkg/controller"
	repo "movie-watchlist/pkg/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	if !repo.InitDb() {
		fmt.Println("Error while initializing database!")
		return
	}

	router := gin.Default()

	router.GET("/movies", controller.GetMovies)

	router.GET("/movies/:id", controller.GetMovie)

	router.GET("/movies/delete/:id", controller.DeleteMovie)

	router.GET("/movies/insert", controller.InsertMovie)

	router.GET("/movies/rate/:id", controller.RateMovie)

	router.Run("localhost:9800")

}
