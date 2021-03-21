package mysql

import (
	"database/sql"

	"github.com/ssentinull/stockbit-assignment/pkg/domain"
)

type movieMySQLRepository struct {
	db *sql.DB
}

func NewMovieMySQLRepository(dummyDB *sql.DB) domain.MovieMySQLRepository {
	return &movieMySQLRepository{
		db: dummyDB,
	}
}
