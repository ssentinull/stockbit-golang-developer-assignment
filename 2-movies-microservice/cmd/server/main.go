package main

import (
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/stockbit-assignment/config"
	"github.com/ssentinull/stockbit-assignment/db"
	_movieHttpHndlr "github.com/ssentinull/stockbit-assignment/pkg/movie/handler/http"
	_movieMySQLRepo "github.com/ssentinull/stockbit-assignment/pkg/movie/repository/mysql"
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
	mysqlDBConn := db.NewMySQLDBConn()
	movieMysqlRepo := _movieMySQLRepo.NewMovieMySQLRepository(mysqlDBConn)
	movieOMDBRepo := _movieOMDBRepo.NewMovieOMDBRepository()
	movieUsecase := _movieUcase.NewMovieUsecase(movieOMDBRepo, movieMysqlRepo)
	_movieHttpHndlr.NewMovieHttpHandler(e, movieUsecase)

	s := &http.Server{
		Addr: config.ServerPort(),
	}

	logrus.Fatal(e.StartServer(s))
}
