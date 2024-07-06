package db

import (
	"context"
	"fmt"
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/utils"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func InitDb() {
	checkEnvVars()
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
		Addr:     os.Getenv("PSQL-ADDRESS-PROD"),
		User:     os.Getenv("PSQL-USER"),
		Password: os.Getenv("PSQL-PASS"),
		Database: os.Getenv("PSQL-DB"),
	})
}

func checkEnvVars() {
	fmt.Println("====== Checking environment ======")
	fmt.Println(os.Getenv("PSQL-ADDRESS-PROD"))
	fmt.Println(os.Getenv("PSQL-USER"))
	fmt.Println(os.Getenv("PSQL-PASS"))
	fmt.Println(os.Getenv("PSQL-DB"))
}
