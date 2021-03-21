package usecase

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/ssentinull/stockbit-assignment/pkg/domain"
	"github.com/ssentinull/stockbit-assignment/pkg/utils"
	httpUtils "github.com/ssentinull/stockbit-assignment/pkg/utils/http"
)

type movieUsecase struct {
	movieOMDBRepository  domain.MovieOMDBRepository
	movieMySQLRepository domain.MovieMySQLRepository
}

func NewMovieUsecase(mor domain.MovieOMDBRepository, mmr domain.MovieMySQLRepository) domain.MovieUsecase {
	return &movieUsecase{
		movieOMDBRepository:  mor,
		movieMySQLRepository: mmr,
	}
}

func (mu *movieUsecase) GetMovieByTitle(ctx context.Context, csr *httpUtils.Cursor) (domain.MovieDetails, error) {
	logger := logrus.WithFields(logrus.Fields{
		"context": utils.Dump(ctx),
		"cursor":  utils.Dump(csr),
	})

	movie, err := mu.movieOMDBRepository.ReadMovieByTitle(ctx, csr)
	if err != nil {
		logger.Error(err)

		return domain.MovieDetails{}, err
	}

	return movie, nil
}

func (mu *movieUsecase) GetMovies(ctx context.Context, csr *httpUtils.Cursor) ([]domain.Movie, error) {
	logger := logrus.WithFields(logrus.Fields{
		"context": utils.Dump(ctx),
		"cursor":  utils.Dump(csr),
	})

	movies, err := mu.movieOMDBRepository.ReadMovies(ctx, csr)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	err = mu.movieMySQLRepository.CreateSearchMovieLog(ctx, csr)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	return movies, nil
}
