package test

import (
	"bytes"
	"encoding/json"
	"io"
	"movie-watchlist/pkg/controller"
	"movie-watchlist/pkg/models"
	"movie-watchlist/pkg/utils"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var base = "/movies"

type MovieDTO struct {
	Name string `json:"name"`
}

func TestAddMovieMustAddMovie(t *testing.T) {
	url := base + "/insert"
	router := setUpRouter()
	router.POST(url, controller.AddMovie)

	movie := MovieDTO{Name: "test-movie"}
	json, err := json.Marshal(movie)

	if utils.CheckError(err) {
		t.Fatalf("Error marshalling body to json")
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(json))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	responseData, _ := io.ReadAll(recorder.Body)

	response := string(responseData)
	result, _ := strconv.ParseBool(response)

	assert.True(t, result)
}

func TestAddMovieMustNotAddMovieWithoutName(t *testing.T) {
	url := base + "/insert"
	router := setUpRouter()
	router.POST(url, controller.AddMovie)

	movie := MovieDTO{}
	json, err := json.Marshal(movie)

	if utils.CheckError(err) {
		t.Fatalf("Error marshalling body to json")
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(json))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	responseData, _ := io.ReadAll(recorder.Body)

	response := string(responseData)
	result, _ := strconv.ParseBool(response)

	assert.False(t, result)
}

func TestGetMoviesMustReturnMovies(t *testing.T) {
	router := setUpRouter()
	router.GET(base, controller.GetMovies)

	req, _ := http.NewRequest("GET", "/movies", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	responseData, _ := io.ReadAll(recorder.Body)
	response := string(responseData)

	var movies []models.Movie

	err := json.Unmarshal([]byte(response), &movies)

	if utils.CheckError(err) {
		t.Fatal("Error unmarshalling Response")
	}

	if len(movies) == 0 {
		t.Fatal("Error unmarshalling Response")
	}
}

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
