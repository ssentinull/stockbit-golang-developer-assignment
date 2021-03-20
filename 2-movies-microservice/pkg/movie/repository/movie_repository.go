package repository

import (
	"context"

	"github.com/ssentinull/stockbit-assignment/pkg/domain"
)

type movieRepository struct{}

func NewMovieRepository() domain.MovieRepository {
	return &movieRepository{}
}

func (mr *movieRepository) ReadMovies(ctx context.Context) string {
	return "read movies"
}
