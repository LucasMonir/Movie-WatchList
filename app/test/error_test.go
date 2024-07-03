package test

import (
	"fmt"
	"movie-watchlist/pkg/utils"
	"testing"
)

func TestCheckErrorWithError(t *testing.T) {
	err := fmt.Errorf("TEST MESSAGE: error while opening database")

	if !utils.CheckError(err) {
		t.Fatal("Expected CheckError to return true, it returned false")
	}
}

func TestCheckErrorWithoutError(t *testing.T) {
	if utils.CheckError(nil) {
		t.Fatal("Expected CheckError to return false, it returned true")
	}
}
