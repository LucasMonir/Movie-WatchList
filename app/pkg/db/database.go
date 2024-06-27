package db

import (
	"context"
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/utils"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func InitDb() {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	if utils.CheckError(db.Ping(ctx)) {
		return
	}

	err := CreateSchema(db)

	if utils.CheckError(err) {
		return
	}

}

func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.Movie)(nil),
	}

	for _, model := range models {

		err := db.Model(model).CreateTable(&orm.CreateTableOptions{IfNotExists: true})

		if utils.CheckError(err) {
			return err
		}
	}

	return nil
}

func GetConnection() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "lu123cas",
		Database: "movie-watchlist",
	})
}
