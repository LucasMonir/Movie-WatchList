package models

type Movie struct {
	Id      int     `json:"id"`
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
