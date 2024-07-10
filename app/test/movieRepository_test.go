package test

import (
	"movie-watchlist/pkg/repository"
	"movie-watchlist/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMovieMustReturnTrueAndId(t *testing.T) {
	name := "Movie For Testing"

	result, id, err := repository.CreateMovie(name)

	if utils.CheckError(err) {
		t.Fatal("No error was expected to happen")
	}

	if id == 0 {
		t.Fatal("Id should return a number above 0")
	}

	assert.True(t, result, "Error while creating movie")
}

func TestCreateMovieWithoutNameMustReturnFalse(t *testing.T) {
	result, id, err := repository.CreateMovie("")

	if !utils.CheckError(err) {
		t.Fatal("An error was expected to happen")
	}

	if id != 0 {
		t.Fatal("Id should return 0")
	}

	assert.False(t, result, "Movie shouldn't be created")
}
