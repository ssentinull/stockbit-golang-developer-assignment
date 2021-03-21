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

type MovieDetails struct {
	Title             string   `json:"Title"`
	ReleaseYear       string   `json:"Year"`
	AgeRating         string   `json:"Rated"`
	ReleaseDate       string   `json:"Released"`
	Runtime           string   `json:"Runtime"`
	Genre             string   `json:"Genre"`
	Director          string   `json:"Director"`
	Writer            string   `json:"Writer"`
	Actors            string   `json:"Actors"`
	Plot              string   `json:"Plot"`
	Language          string   `json:"Language"`
	Country           string   `json:"Country"`
	Awards            string   `json:"Awards"`
	Poster            string   `json:"Poster"`
	Ratings           []Rating `json:"Ratings"`
	Metascore         string   `json:"Metascore"`
	IMDBRating        string   `json:"imdbRating"`
	IMDBVotes         string   `json:"imdbVotes"`
	IMDBID            string   `json:"imdbID"`
	Type              string   `json:"Type"`
	DVDReleaseDate    string   `json:"DVD"`
	BoxOffice         string   `json:"BoxOffice"`
	ProductionCompany string   `json:"Production"`
	Website           string   `json:"Website"`
	ErrorMessage      string   `json:"Error"`
}

type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

type MovieUsecase interface {
	GetMovieByTitle(context.Context, *httpUtils.Cursor) (MovieDetails, error)
	GetMovies(context.Context, *httpUtils.Cursor) ([]Movie, error)
}

type MovieOMDBRepository interface {
	ReadMovieByTitle(context.Context, *httpUtils.Cursor) (MovieDetails, error)
	ReadMovies(context.Context, *httpUtils.Cursor) ([]Movie, error)
}

type MovieMySQLRepository interface {
	CreateGetMoviesLog(context.Context, *httpUtils.Cursor) error
}
