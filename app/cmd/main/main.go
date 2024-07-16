package main

import (
	"movie-watchlist/pkg/controller"
	"movie-watchlist/pkg/db"
	"movie-watchlist/pkg/queue"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	startServer()
}

func startServer() {
	base := "/movies"
	router := gin.Default()
	router.GET(base, controller.GetMovies)
	router.GET(base+"/:id", controller.GetMovie)
	router.POST(base+"/insert", controller.AddMovie)
	router.PATCH(base+"/rate/:id", controller.RateMovie)
	router.DELETE(base+"/delete/:id", controller.DeleteMovie)

	router.Run("0.0.0.0:9800")
	queue.SendLogToServer("App server started: port 9800!")
}
