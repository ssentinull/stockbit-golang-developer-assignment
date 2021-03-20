package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ssentinull/stockbit-assignment/pkg/domain"
)

type movieHttpHandler struct {
	movieUsecase domain.MovieUsecase
}

func NewMovieHttpHandler(e *echo.Echo, mu domain.MovieUsecase) {
	handler := &movieHttpHandler{
		movieUsecase: mu,
	}

	api := e.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/movies", handler.FetchMovies)
}

func (mhh *movieHttpHandler) FetchMovies(c echo.Context) error {
	movies := mhh.movieUsecase.GetMovies(c.Request().Context())

	return c.JSON(http.StatusOK, movies)
}
