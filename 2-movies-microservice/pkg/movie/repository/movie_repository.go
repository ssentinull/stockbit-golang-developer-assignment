package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/ssentinull/stockbit-assignment/config"
	"github.com/ssentinull/stockbit-assignment/pkg/domain"
)

type movieRepository struct{}

func NewMovieRepository() domain.MovieRepository {
	return &movieRepository{}
}

func (mr *movieRepository) ReadMovies(ctx context.Context) ([]domain.Movie, error) {
	logger := logrus.WithField("context", ctx)

	omdbAPIBaseURL := "http://www.omdbapi.com/?"
	omdbAPIKey := config.OMDBKey()
	requestURL := fmt.Sprintf("%sapikey=%s&s=Batman", omdbAPIBaseURL, omdbAPIKey)
	resp, err := http.Get(requestURL)
	if err != nil {
		logger.Error(err)

		return nil, err
	}
	defer resp.Body.Close()

	omdbRes := new(domain.OMDBSearchResponse)
	err = json.NewDecoder(resp.Body).Decode(omdbRes)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	return omdbRes.Payload, nil
}
