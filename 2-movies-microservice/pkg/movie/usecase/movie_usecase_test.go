package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ssentinull/stockbit-assignment/pkg/domain"
	customMock "github.com/ssentinull/stockbit-assignment/pkg/domain/mock"
	"github.com/ssentinull/stockbit-assignment/pkg/movie/usecase"
	httpUtils "github.com/ssentinull/stockbit-assignment/pkg/utils/http"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetMovies(t *testing.T) {
	mockMovie := domain.Movie{
		Title:  "Avatar",
		Year:   "2010",
		IMDBID: "tt0499549",
		Type:   "movie",
		Poster: "https://m.media-amazon.com/images/M/MV5BMTYwOTEwNjAzMl5BMl5BanBnXkFtZTcwODc5MTUwMw@@._V1_SX300.jpg",
	}
	cursor := httpUtils.Cursor{}

	mockMovies := make([]domain.Movie, 0)
	mockMovies = append(mockMovies, mockMovie)
	mockOMDBRepo := new(customMock.MovieOMDBRepository)
	mockMySQLRepo := new(customMock.MovieMySQLRepository)
	t.Run("success", func(t *testing.T) {
		mockOMDBRepo.On("ReadMovies", mock.Anything, mock.Anything).
			Return(mockMovies, nil).Once()
		mockMySQLRepo.On("CreateGetMoviesLog", mock.Anything, mock.Anything).
			Return(nil).Once()

		u := usecase.NewMovieUsecase(mockOMDBRepo, mockMySQLRepo)
		movies, err := u.GetMovies(context.TODO(), &cursor)

		assert.NoError(t, err)
		assert.NotNil(t, movies)
		assert.Len(t, mockMovies, len(movies))

		mockOMDBRepo.AssertExpectations(t)
		mockMySQLRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockOMDBRepo.On("ReadMovies", mock.Anything, mock.Anything).
			Return(nil, errors.New("error")).Once()

		u := usecase.NewMovieUsecase(mockOMDBRepo, mockMySQLRepo)
		movies, err := u.GetMovies(context.TODO(), &cursor)

		assert.Error(t, err)
		assert.Len(t, movies, 0)

		mockOMDBRepo.AssertExpectations(t)
		mockMySQLRepo.AssertExpectations(t)
	})
}
