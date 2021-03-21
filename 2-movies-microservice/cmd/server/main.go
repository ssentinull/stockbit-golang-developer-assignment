package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/stockbit-assignment/config"
	_movieHttpHndlr "github.com/ssentinull/stockbit-assignment/pkg/movie/handler/http"
	_movieOMDBRepo "github.com/ssentinull/stockbit-assignment/pkg/movie/repository/omdb"
	_movieUcase "github.com/ssentinull/stockbit-assignment/pkg/movie/usecase"
)

func initLogger() {
	logLevel := logrus.ErrorLevel
	switch config.Env() {
	case "development":
		logLevel = logrus.InfoLevel
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:    true,
		DisableSorting: true,
		DisableColors:  false,
	})

	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
	logrus.SetLevel(logLevel)
}

func init() {
	initLogger()
}

func main() {
	e := echo.New()
	movieOMDBRepo := _movieOMDBRepo.NewMovieOMDBRepository()
	movieUsecase := _movieUcase.NewMovieUsecase(movieOMDBRepo)
	_movieHttpHndlr.NewMovieHttpHandler(e, movieUsecase)

	s := &http.Server{
		Addr: config.ServerPort(),
	}

	logrus.Fatal(e.StartServer(s))
}
