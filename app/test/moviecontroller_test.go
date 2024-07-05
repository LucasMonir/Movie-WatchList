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

type RatingDTO struct {
	Rating float32 `json:"rating"`
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

	assert.False(t, result, "Fail: Result must be 'False'")
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

	if utils.CheckError(err) && recorder.Result().StatusCode != 200 {
		t.Fatal("Error unmarshalling Response")
	}
}

func TestRateMovieShouldAddRating(t *testing.T) {
	url := base + "/rate/:id"
	router := setUpRouter()
	router.PATCH(url, controller.RateMovie)

	ratingParam := "/rate/1?rating=10"

	req, err := http.NewRequest("PATCH", base+ratingParam, nil)
	recorder := httptest.NewRecorder()

	if utils.CheckError(err) {
		t.Fatal("Error during the request")
	}

	router.ServeHTTP(recorder, req)

	response, _ := io.ReadAll(recorder.Body)
	responseData := string(response)
	result, _ := strconv.ParseBool(responseData)

	assert.True(t, result, "Fail: result must be 'True'")
}

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
