package domain

import (
	"context"

	"github.com/ssentinull/stockbit-assignment/pkg/utils"
	httpUtils "github.com/ssentinull/stockbit-assignment/pkg/utils/http"
)

type OMDBSearchMoviesResponse struct {
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

type OMDBGetMovieResponse struct {
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

type MovieDetails struct {
	Title               string
	ReleaseYear         string
	AgeRating           string
	ReleaseDate         string
	Runtime             string
	Genre               []string
	Directors           []string
	Writers             []string
	Actors              []string
	Plot                string
	Languages           []string
	Countries           []string
	Awards              string
	Poster              string
	Ratings             []Rating
	Metascore           string
	IMDBRating          string
	IMDBVotes           string
	IMDBID              string
	Type                string
	DVDReleaseDate      string
	BoxOffice           string
	ProductionCompanies []string
	Website             string
	ErrorMessage        string
}

func OMDBResponseToMovieDetails(omdbRes OMDBGetMovieResponse) MovieDetails {
	genres := utils.SplitStringByComma(omdbRes.Genre)
	directors := utils.SplitStringByComma(omdbRes.Director)
	writers := utils.SplitStringByComma(omdbRes.Writer)
	actors := utils.SplitStringByComma(omdbRes.Actors)
	languages := utils.SplitStringByComma(omdbRes.Language)
	countries := utils.SplitStringByComma(omdbRes.Country)
	productionCompanies := utils.SplitStringByComma(omdbRes.ProductionCompany)

	return MovieDetails{
		Title:               omdbRes.Title,
		ReleaseYear:         omdbRes.ReleaseYear,
		AgeRating:           omdbRes.AgeRating,
		ReleaseDate:         omdbRes.ReleaseDate,
		Runtime:             omdbRes.Runtime,
		Genre:               genres,
		Directors:           directors,
		Writers:             writers,
		Actors:              actors,
		Plot:                omdbRes.Plot,
		Languages:           languages,
		Countries:           countries,
		Awards:              omdbRes.Awards,
		Poster:              omdbRes.Poster,
		Ratings:             omdbRes.Ratings,
		Metascore:           omdbRes.Metascore,
		IMDBRating:          omdbRes.IMDBRating,
		IMDBVotes:           omdbRes.IMDBVotes,
		IMDBID:              omdbRes.IMDBID,
		Type:                omdbRes.Type,
		DVDReleaseDate:      omdbRes.DVDReleaseDate,
		BoxOffice:           omdbRes.BoxOffice,
		ProductionCompanies: productionCompanies,
		Website:             omdbRes.Website,
		ErrorMessage:        omdbRes.ErrorMessage,
	}
}

type MovieUsecase interface {
	GetMovieByTitle(context.Context, *httpUtils.Cursor) (MovieDetails, error)
	GetMovies(context.Context, *httpUtils.Cursor) ([]Movie, error)
}

type MovieOMDBRepository interface {
	ReadMovieByTitle(context.Context, *httpUtils.Cursor) (OMDBGetMovieResponse, error)
	ReadMovies(context.Context, *httpUtils.Cursor) ([]Movie, error)
}

type MovieMySQLRepository interface {
	CreateGetMovieByTitleLog(context.Context, *httpUtils.Cursor) error
	CreateGetMoviesLog(context.Context, *httpUtils.Cursor) error
}
