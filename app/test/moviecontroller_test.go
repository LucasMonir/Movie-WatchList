package test

import (
	"bytes"
	"encoding/json"
	"fmt"
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
var movieId = 0

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
	result, _ := strconv.Atoi(response)

	assert.NotEqual(t, movieId, result)
	movieId = result
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

	ratingParam := fmt.Sprintf("/rate/%d?rating=10", movieId)

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

func TestRateMovieShouldNotAddRating(t *testing.T) {
	url := base + "/rate/:id"
	router := setUpRouter()
	router.PATCH(url, controller.RateMovie)

	ratingParam := fmt.Sprintf("/rate/%d?rating=14", movieId)

	req, err := http.NewRequest("PATCH", base+ratingParam, nil)
	recorder := httptest.NewRecorder()

	if utils.CheckError(err) {
		t.Fatal("Error during the request")
	}

	router.ServeHTTP(recorder, req)

	assert.NotEqual(t, int(http.StatusOK), recorder.Result().StatusCode)
}

func TestDeleteMovieShouldDeleteMovie(t *testing.T) {
	url := base + "/delete/:id"
	deleteUrl := base + fmt.Sprintf("/delete/%d", movieId)

	router := setUpRouter()
	router.DELETE(url, controller.DeleteMovie)

	req, err := http.NewRequest("DELETE", deleteUrl, nil)
	recorder := httptest.NewRecorder()

	if utils.CheckError(err) {
		t.Fatal("Error during the request")
	}

	router.ServeHTTP(recorder, req)

	response, _ := io.ReadAll(recorder.Body)
	responseData := string(response)
	result, _ := strconv.ParseBool(responseData)

	assert.True(t, result)
}

func TestDeleteMovieShouldNotDeleteMovieInvalidId(t *testing.T) {
	url := base + "/delete/:id"
	deleteUrl := base + fmt.Sprintf("/delete/%d", fakeId(movieId))

	router := setUpRouter()
	router.DELETE(url, controller.DeleteMovie)

	req, err := http.NewRequest("DELETE", deleteUrl, nil)
	recorder := httptest.NewRecorder()

	if utils.CheckError(err) {
		t.Fatal("Error during the request")
	}

	router.ServeHTTP(recorder, req)

	response, _ := io.ReadAll(recorder.Body)
	responseData := string(response)
	result, _ := strconv.ParseBool(responseData)

	assert.False(t, result)
}

func setUpRouter() *gin.Engine {
	router := gin.New()
	return router
}

func fakeId(id int) int {
	return id * 17
}
