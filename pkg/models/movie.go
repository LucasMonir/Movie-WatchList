package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name    string  `json:"name"`
	Watched bool    `json:"watched"`
	Rating  float32 `json:"rating"`
}

func NewMovie(name string) Movie {
	return Movie{
		Name:    name,
		Watched: false,
		Rating:  0,
	}
}
