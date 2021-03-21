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

var ()

type movieMySQLRepository struct {
	db *sql.DB
}

func NewMovieMySQLRepository(dummyDB *sql.DB) domain.MovieMySQLRepository {
	return &movieMySQLRepository{
		db: dummyDB,
	}
}

func (mmr *movieMySQLRepository) CreateGetMovieByTitleLog(ctx context.Context, csr *httpUtils.Cursor) error {
	logger := logrus.WithFields(logrus.Fields{
		"context": utils.Dump(ctx),
		"cursor":  utils.Dump(csr),
	})

	args := make([]interface{}, 0)
	arg := genCreateGetMovieByTitleLogArg(csr)
	jakartaTime, err := mmr.getJakartaCurrentTime()
	if err != nil {
		logger.Error(err)

		return err
	}

	args = append(args, arg, jakartaTime)
	query := genCreateGetMovieByTitleLogQry()
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

func (mmr *movieMySQLRepository) CreateGetMoviesLog(ctx context.Context, csr *httpUtils.Cursor) error {
	logger := logrus.WithFields(logrus.Fields{
		"context": utils.Dump(ctx),
		"cursor":  utils.Dump(csr),
	})

	args := genCreateGetMoviesLogArgs(csr)
	jakartaTime, err := mmr.getJakartaCurrentTime()
	if err != nil {
		logger.Error(err)

		return err
	}

	args = append(args, jakartaTime)
	query := genCreateGetMoviesLogQry()
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

func (mmr *movieMySQLRepository) getJakartaCurrentTime() (time.Time, error) {
	jkt, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return time.Time{}, err
	}

	return time.Now().In(jkt), nil
}
