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

func MySQLDBDSN() string {
	dbUser := os.Getenv("MYSQL_DB_USER")
	dbPassword := os.Getenv("MYSQL_DB_PASSWORD")
	dbProtocol := os.Getenv("MYSQL_DB_PROTOCOL")
	dbHost := os.Getenv("MYSQL_DB_HOST")
	dbPort := os.Getenv("MYSQL_DB_PORT")
	dbName := os.Getenv("MYSQL_DB_NAME")

	return fmt.Sprintf("%v:%v@%v(%v:%v)/%v?parseTime=true",
		dbUser, dbPassword, dbProtocol, dbHost, dbPort, dbName)
}

func OMDBKey() string {
	val, ok := os.LookupEnv("OMDB_API_KEY")
	if !ok {
		logrus.Fatal("OMDB_API_KEY not provided")
	}

	return val
}

func ServerPort() string {
	return fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
}
