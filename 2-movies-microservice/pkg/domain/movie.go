package domain

import "context"

type Movie struct{}

type MovieUsecase interface {
	GetMovies(context.Context) string
}

type MovieRepository interface {
	ReadMovies(context.Context) string
}
