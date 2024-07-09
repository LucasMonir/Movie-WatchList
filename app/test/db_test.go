package test

import (
	"movie-watchlist/pkg/db"
	"movie-watchlist/pkg/utils"
	"testing"
)

func TestGetConnectionMustReturnConnection(t *testing.T) {
	db := db.GetConnection()
	defer db.Close()

	err := db.Ping(db.Context())

	if utils.CheckError(err) {
		t.Fatal("Database connection not created")
	}
}
