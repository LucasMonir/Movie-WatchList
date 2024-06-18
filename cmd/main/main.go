package main

import (
	"fmt"
	"movie-watchlist/pkg/controller"
	repo "movie-watchlist/pkg/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	if !repo.CheckStorageCreated() {
		fmt.Println("Error while initializing database!")
		return
	}

	startServer()
}

func startServer() {
	base := "/movies"
	router := gin.Default()

	router.GET(base, controller.GetMovies)
	router.GET(base+"/:id", controller.GetMovie)
	router.DELETE(base+"/delete/:id", controller.DeleteMovie)
	router.POST(base+"/insert", controller.AddMovie)
	router.PATCH(base+"/rate/:id", controller.RateMovie)
	router.Run("localhost:9800")
}
