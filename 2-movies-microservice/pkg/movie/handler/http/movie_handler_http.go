package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/stockbit-assignment/pkg/domain"
	httpUtils "github.com/ssentinull/stockbit-assignment/pkg/utils/http"
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
	v1.GET("/movies/:title", handler.FetchMovieByTitle)
}

func (mhh *movieHttpHandler) FetchMovies(c echo.Context) error {
	cursor, err := httpUtils.NewCursor(c)
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, err.Error())
	}

	movies, err := mhh.movieUsecase.GetMovies(c.Request().Context(), cursor)
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	movieResponses := newMovieResponses(movies)

	return c.JSON(http.StatusOK, httpUtils.NewCursorResponse(cursor, movieResponses))
}

func (mhh *movieHttpHandler) FetchMovieByTitle(c echo.Context) error {
	cursor, err := httpUtils.NewCursor(c)
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, err.Error())
	}

	movie, err := mhh.movieUsecase.GetMovieByTitle(c.Request().Context(), cursor)
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	movieDetailsRes := newMovieDetailsResponse(movie)

	return c.JSON(http.StatusOK, httpUtils.NewCursorResponse(cursor, movieDetailsRes))
}
