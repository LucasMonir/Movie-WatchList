package models

import "fmt"

type Movie struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Watched bool    `json:"watched"`
	Rating  float32 `json:"rating"`
}

func NewMovie(name string) (Movie, error) {
	movie := Movie{}

	if !checkMovie(name) {
		return movie, fmt.Errorf("invalid name parameter")
	}

	movie = Movie{
		Name:    name,
		Watched: false,
		Rating:  0,
	}

	return movie, nil
}

func checkMovie(name string) bool {
	return len(name) > 0 && name != "" && name != " "
}
