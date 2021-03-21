package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/ssentinull/stockbit-assignment/pkg/domain"
	"github.com/ssentinull/stockbit-assignment/pkg/utils"
	httpUtils "github.com/ssentinull/stockbit-assignment/pkg/utils/http"
)

type movieMySQLRepository struct {
	db *sql.DB
}

func NewMovieMySQLRepository(dummyDB *sql.DB) domain.MovieMySQLRepository {
	return &movieMySQLRepository{
		db: dummyDB,
	}
}

func (mmr *movieMySQLRepository) CreateSearchMovieLog(ctx context.Context, csr *httpUtils.Cursor) error {
	logger := logrus.WithFields(logrus.Fields{
		"context": utils.Dump(ctx),
		"cursor":  utils.Dump(csr),
	})

	args := genCreateSearchLogArgs(csr)
	jkt, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		logger.Error(err)

		return err
	}

	args = append(args, time.Now().In(jkt))
	query := genCreateSearchLogQry()
	stmt, err := mmr.db.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(err)

		return err
	}

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		logger.Error(err)

		return err
	}

	return nil
}
