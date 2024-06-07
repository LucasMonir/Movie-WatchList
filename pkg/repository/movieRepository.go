package repository

import (
	models "movie-watchlist/pkg/models"
	utils "movie-watchlist/pkg/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDb() {
	db, err := gorm.Open(sqlite.Open("movies.db"), &gorm.Config{})

	utils.CheckError(err)

	db.AutoMigrate(&models.Movie{})
}
