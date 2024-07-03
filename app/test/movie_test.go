package test

import (
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/utils"
	"reflect"
	"testing"
)

func TestNewMovieWithName(t *testing.T) {
	name := "Movie-01"
	emptyMovie := models.Movie{}
	movie, err := models.NewMovie(name)

	if utils.CheckError(err) {
		t.Fatal(reflect.ValueOf(TestNewMovieWithName).Pointer(), "No error was expected")
	}

	if movie.Name == emptyMovie.Name {
		t.Fatal(reflect.ValueOf(TestNewMovieWithName).Pointer(), "No movie was created")
	}

	if movie.Name != name {
		t.Fatal(reflect.ValueOf(TestNewMovieWithName).Pointer(), "Movie incorrectly created")
	}
}

func TestNewMovieWithoutName(t *testing.T) {
	name := ""
	emptyMovie := models.Movie{}
	movie, err := models.NewMovie(name)

	if !utils.CheckError(err) {
		t.Fatal(reflect.ValueOf(TestNewMovieWithName).Pointer(), "Expected error, none given")
	}

	if movie.Name != emptyMovie.Name {
		t.Fatal(reflect.ValueOf(TestNewMovieWithName).Pointer(), "Movie shouldn't be created")
	}
}
