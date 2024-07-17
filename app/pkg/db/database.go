package db

import (
	"context"
	"fmt"
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/queue"
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

	queue.SendLogToServer("Database initialized successfully!")
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
		Addr:     os.Getenv("PSQL_ADDRESS_PROD"),
		User:     os.Getenv("PSQL_USER"),
		Password: os.Getenv("PSQL_PASS"),
		Database: os.Getenv("PSQL_DB"),
	})
}

func checkEnvVars() {
	fmt.Println("====== Checking environment ======")
	fmt.Println(os.Getenv("PSQL_ADDRESS_PROD"))
	fmt.Println(os.Getenv("PSQL_USER"))
	fmt.Println(os.Getenv("PSQL_PASS"))
	fmt.Println(os.Getenv("PSQL_DB"))
	fmt.Println(os.Getenv("RABBIT_MQ_PROD"))
	fmt.Println("====== Finished environment check ======")
}
