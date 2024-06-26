package db

import (
	"context"
	"fmt"
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/utils"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func InitDb() {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "lu123cas",
		Database: "movie-watchlist",
	})

	defer db.Close()

	ctx := context.Background()

	if utils.CheckError(db.Ping(ctx)) {
		return
	}

	fmt.Println("db created, creating schema")

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
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{})

		if utils.CheckError(err) {
			return err
		}
	}

	return nil
}
