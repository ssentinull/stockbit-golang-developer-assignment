package usecase

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/ssentinull/stockbit-assignment/pkg/domain"
)

type movieUsecase struct {
	movieRepository domain.MovieRepository
}

func NewMovieUsecase(mr domain.MovieRepository) domain.MovieUsecase {
	return &movieUsecase{
		movieRepository: mr,
	}
}

func (mu *movieUsecase) GetMovies(ctx context.Context) ([]domain.Movie, error) {
	movies, err := mu.movieRepository.ReadMovies(ctx)
	if err != nil {
		logrus.WithField("context", ctx).Error(err)

		return nil, err
	}

	return movies, nil
}
