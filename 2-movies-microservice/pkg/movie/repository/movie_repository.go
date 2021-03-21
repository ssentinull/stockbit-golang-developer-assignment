package repository

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

type movieRepository struct{}

func NewMovieRepository() domain.MovieRepository {
	return &movieRepository{}
}

func (mr *movieRepository) ReadMovies(ctx context.Context, csr *httpUtils.Cursor) ([]domain.Movie, error) {
	logger := logrus.WithFields(logrus.Fields{
		"context": utils.Dump(ctx),
		"cursor":  utils.Dump(csr),
	})

	omdbAPIBaseURL := "http://www.omdbapi.com/?"
	omdbAPIKey := config.OMDBKey()
	requestURL := genReadMoviesURL(omdbAPIBaseURL, omdbAPIKey, csr)
	requestRes, err := http.Get(requestURL)
	if err != nil {
		logger.Error(err)

		return nil, err
	}
	defer requestRes.Body.Close()

	omdbRes := new(domain.OMDBSearchResponse)
	err = json.NewDecoder(requestRes.Body).Decode(omdbRes)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	return omdbRes.Payload, nil
}
