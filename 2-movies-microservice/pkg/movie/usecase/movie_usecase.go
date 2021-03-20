package usecase

import (
	"context"

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

func (mu *movieUsecase) GetMovies(ctx context.Context) string {
	movies := mu.movieRepository.ReadMovies(ctx)

	return "get movies -> " + movies
}
