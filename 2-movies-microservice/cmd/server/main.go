package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/stockbit-assignment/config"
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
	s := &http.Server{
		Addr: config.ServerPort(),
	}

	log.Fatal(e.StartServer(s))
}
