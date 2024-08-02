package utils

import (
	"os"

	"github.com/go-pg/pg/v10"
)

func GetCustomConnection(options *pg.Options) *pg.DB {
	var emptyOpt *pg.Options

	if options != emptyOpt {
		return pg.Connect(options)
	}

	defaultConfig := getConnectionConfig()
	return pg.Connect(defaultConfig)
}

func getConnectionConfig() *pg.Options {
	return &pg.Options{
		Addr:     os.Getenv("PSQL_ADDRESS_PROD"),
		User:     os.Getenv("PSQL_USER"),
		Password: os.Getenv("PSQL_PASS"),
		Database: os.Getenv("PSQL_DB"),
	}
}
