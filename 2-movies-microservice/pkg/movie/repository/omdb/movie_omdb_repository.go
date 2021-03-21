package omdb

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/ssentinull/stockbit-assignment/config"
	"github.com/ssentinull/stockbit-assignment/pkg/domain"
	"github.com/ssentinull/stockbit-assignment/pkg/utils"
	httpUtils "github.com/ssentinull/stockbit-assignment/pkg/utils/http"
)

var (
	omdbAPIBaseURL = "http://www.omdbapi.com/?"
	omdbAPIKey     = config.OMDBKey()
)

type movieOMDBRepository struct{}

func NewMovieOMDBRepository() domain.MovieOMDBRepository {
	return &movieOMDBRepository{}
}

func (mr *movieOMDBRepository) ReadMovieByTitle(ctx context.Context, csr *httpUtils.Cursor) (domain.OMDBGetMovieResponse, error) {
	logger := logrus.WithFields(logrus.Fields{
		"context": utils.Dump(ctx),
		"cursor":  utils.Dump(csr),
	})

	movieDetails := new(domain.OMDBGetMovieResponse)
	requestURL := genReadMovieByTitle(omdbAPIBaseURL, omdbAPIKey, csr)
	requestRes, err := http.Get(requestURL)
	if err != nil {
		logger.Error(err)
		return *movieDetails, err
	}
	defer requestRes.Body.Close()

	err = json.NewDecoder(requestRes.Body).Decode(movieDetails)
	if err != nil {
		logger.Error(err)

		return *movieDetails, err
	}

	return *movieDetails, nil
}

func (mr *movieOMDBRepository) ReadMovies(ctx context.Context, csr *httpUtils.Cursor) ([]domain.Movie, error) {
	logger := logrus.WithFields(logrus.Fields{
		"context": utils.Dump(ctx),
		"cursor":  utils.Dump(csr),
	})

	requestURL := genReadMoviesURL(omdbAPIBaseURL, omdbAPIKey, csr)
	requestRes, err := http.Get(requestURL)
	if err != nil {
		logger.Error(err)

		return nil, err
	}
	defer requestRes.Body.Close()

	omdbRes := new(domain.OMDBSearchMoviesResponse)
	err = json.NewDecoder(requestRes.Body).Decode(omdbRes)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	return omdbRes.Payload, nil
}
