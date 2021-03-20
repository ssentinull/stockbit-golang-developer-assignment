package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn(".env file not found")
	}

	return
}

func Env() string {
	if val, ok := os.LookupEnv("ENV"); ok {
		return val
	}

	return "development"
}

func ServerPort() string {
	return fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
}
