package domain

import (
	"context"

	httpUtils "github.com/ssentinull/stockbit-assignment/pkg/utils/http"
)

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
	GetMovies(context.Context, *httpUtils.Cursor) ([]Movie, error)
}

type MovieOMDBRepository interface {
	ReadMovies(context.Context, *httpUtils.Cursor) ([]Movie, error)
}
