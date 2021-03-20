package domain

import "context"

type OMDBSearchResponse struct {
	Payload       []Movie `json:"Search"`
	TotalRowCount string  `json:"totalResults"`
	Status        string  `json:"Response"`
	ErrorMessage  string  `json:"Error"`
}

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	IMDBID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type MovieUsecase interface {
	GetMovies(context.Context) ([]Movie, error)
}

type MovieRepository interface {
	ReadMovies(context.Context) ([]Movie, error)
}
