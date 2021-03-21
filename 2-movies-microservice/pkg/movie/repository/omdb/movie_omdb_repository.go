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

type movieOMDBRepository struct{}

func NewMovieOMDBRepository() domain.MovieOMDBRepository {
	return &movieOMDBRepository{}
}

func (mr *movieOMDBRepository) ReadMovies(ctx context.Context, csr *httpUtils.Cursor) ([]domain.Movie, error) {
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
