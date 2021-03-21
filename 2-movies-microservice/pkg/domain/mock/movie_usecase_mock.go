package mock

import (
	"context"

	"github.com/ssentinull/stockbit-assignment/pkg/domain"
	"github.com/ssentinull/stockbit-assignment/pkg/utils/http"
	"github.com/stretchr/testify/mock"
)

type MovieUsecase struct {
	mock.Mock
}

func (_m *MovieUsecase) GetMovies(ctx context.Context, csr http.Cursor) ([]domain.Movie, error) {
	ret := _m.Called(ctx, csr)

	var r0 []domain.Movie
	if rf, ok := ret.Get(0).(func(context.Context, http.Cursor) []domain.Movie); ok {
		r0 = rf(ctx, csr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, http.Cursor) error); ok {
		r1 = rf(ctx, csr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
