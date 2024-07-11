package test

import (
	"movie-watchlist/pkg/repository"
	"movie-watchlist/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testMovieId = 0

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
	testMovieId = id
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

func TestReadAllMoviesMustReturnStoredMovies(t *testing.T) {
	result, err := repository.ReadMovies()

	if utils.CheckError(err) {
		t.Fatal("No error was expected")
	}

	assert.True(t, len(result) > 0, "No movie was found, at least one should be expected")
}

func TestReadMoviesMustReturnMovie(t *testing.T) {
	result, err := repository.ReadMovie(testMovieId)

	if utils.CheckError(err) {
		t.Fatal("No error was expected")
	}

	assert.Equal(t, result.Id, testMovieId, "Movie ID does not match.")
}

func TestReadMoviesMustNotReturnMovieWithInvalidId(t *testing.T) {
	result, err := repository.ReadMovie(fakeId(testMovieId))

	assert.NotNil(t, err, "Error should not be nil, expected error reading movie")
	assert.True(t, result.Id == 0, "Error should not be nil, expected error reading movie")
}
