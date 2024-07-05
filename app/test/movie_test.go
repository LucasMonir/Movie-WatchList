package test

import (
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/utils"
	"testing"
)

func TestNewMovieWithNameShouldBeCreated(t *testing.T) {
	name := "Movie-01"
	emptyMovie := models.Movie{}
	movie, err := models.NewMovie(name)

	if utils.CheckError(err) {
		t.Fatal("No error was expected")
	}

	if movie == emptyMovie {
		t.Fatal("No movie was created")
	}

	if movie.Name != name {
		t.Fatal("Movie incorrectly created")
	}
}

func TestNewMovieWithoutNameShouldNotBeCreated(t *testing.T) {
	name := ""
	emptyMovie := models.Movie{}
	movie, err := models.NewMovie(name)

	if !utils.CheckError(err) {
		t.Fatal("Expected error, none given")
	}

	if movie.Name != emptyMovie.Name {
		t.Fatal("Movie shouldn't be created")
	}
}
