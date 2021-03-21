package usecase

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/ssentinull/stockbit-assignment/pkg/domain"
	"github.com/ssentinull/stockbit-assignment/pkg/utils"
	httpUtils "github.com/ssentinull/stockbit-assignment/pkg/utils/http"
)

type movieUsecase struct {
	movieOMDBRepository domain.MovieOMDBRepository
}

func NewMovieUsecase(mor domain.MovieOMDBRepository) domain.MovieUsecase {
	return &movieUsecase{
		movieOMDBRepository: mor,
	}
}

func (mu *movieUsecase) GetMovies(ctx context.Context, csr *httpUtils.Cursor) ([]domain.Movie, error) {
	movies, err := mu.movieOMDBRepository.ReadMovies(ctx, csr)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"context": utils.Dump(ctx),
			"cursor":  utils.Dump(csr),
		}).Error(err)

		return nil, err
	}

	return movies, nil
}
