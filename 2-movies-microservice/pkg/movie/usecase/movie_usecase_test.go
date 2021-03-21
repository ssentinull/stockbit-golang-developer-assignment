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

func TestGetMovieByTitle(t *testing.T) {
	mockMovie := domain.OMDBGetMovieResponse{
		Title:       "Avatar",
		ReleaseYear: "2009",
		AgeRating:   "PG-13",
		ReleaseDate: "18 Dec 2009",
		Runtime:     "162 min",
		Genre:       "Action, Adventure, Fantasy, Sci-Fi",
		Director:    "James Cameron",
		Writer:      "James Cameron",
		Actors:      "Sam Worthington, Zoe Saldana, Sigourney Weaver, Stephen Lang",
		Plot:        "A paraplegic Marine dispatched to the moon Pandora on a unique mission becomes torn between following his orders and protecting the world he feels is his home.",
		Language:    "English, Spanish",
		Country:     "USA",
		Awards:      "Won 3 Oscars. Another 86 wins & 130 nominations.",
		Poster:      "https://m.media-amazon.com/images/M/MV5BMTYwOTEwNjAzMl5BMl5BanBnXkFtZTcwODc5MTUwMw@@._V1_SX300.jpg",
		Ratings: []domain.Rating{
			{
				Source: "Internet Movie Database",
				Value:  "7.8/10",
			},
			{
				Source: "Rotten Tomatoes",
				Value:  "82%",
			},
			{
				Source: "Metacritic",
				Value:  "83/100",
			},
		},
		Metascore:         "83",
		IMDBRating:        "7.8",
		IMDBVotes:         "1,124,863",
		IMDBID:            "tt0499549",
		Type:              "movie",
		DVDReleaseDate:    "10 Feb 2016",
		BoxOffice:         "$760,507,625",
		ProductionCompany: "Dune, Lightstorm Entertainment, Ingenious Film Partners",
		Website:           "N/A",
		ErrorMessage:      "",
	}
	cursor := httpUtils.Cursor{}

	mockOMDBRepo := new(customMock.MovieOMDBRepository)
	mockMySQLRepo := new(customMock.MovieMySQLRepository)
	t.Run("success", func(t *testing.T) {
		mockOMDBRepo.On("ReadMovieByTitle", mock.Anything, mock.Anything).
			Return(mockMovie, nil).Once()
		mockMySQLRepo.On("CreateGetMovieByTitleLog", mock.Anything, mock.Anything).
			Return(nil).Once()

		u := usecase.NewMovieUsecase(mockOMDBRepo, mockMySQLRepo)
		movies, err := u.GetMovieByTitle(context.TODO(), &cursor)

		assert.NoError(t, err)
		assert.NotNil(t, movies)

		mockOMDBRepo.AssertExpectations(t)
		mockMySQLRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockOMDBRepo.On("ReadMovieByTitle", mock.Anything, mock.Anything).
			Return(nil, errors.New("error")).Once()

		u := usecase.NewMovieUsecase(mockOMDBRepo, mockMySQLRepo)
		movie, err := u.GetMovieByTitle(context.TODO(), &cursor)

		assert.Error(t, err)
		assert.Equal(t, domain.MovieDetails{}, movie)

		mockOMDBRepo.AssertExpectations(t)
		mockMySQLRepo.AssertExpectations(t)
	})
}
