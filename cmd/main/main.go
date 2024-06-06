package main

import (
	"movie-watchlist/pkg/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/movies", controller.GetMovies)

	router.GET("/movies/:id", controller.GetMovie)

	router.GET("/movies/delete/:id", controller.DeleteMovie)

	router.GET("/movies/insert", controller.InsertMovie)

	router.GET("/movies/rate/:id", controller.RateMovie)

	router.Run("localhost:9800")

}
