package mock

import (
	"context"

	"github.com/ssentinull/stockbit-assignment/pkg/utils/http"
	"github.com/stretchr/testify/mock"
)

type MovieMySQLRepository struct {
	mock.Mock
}

func (_m *MovieMySQLRepository) CreateGetMovieByTitleLog(ctx context.Context, csr *http.Cursor) error {
	ret := _m.Called(ctx, csr)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *http.Cursor) error); ok {
		r0 = rf(ctx, csr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MovieMySQLRepository) CreateGetMoviesLog(ctx context.Context, csr *http.Cursor) error {
	ret := _m.Called(ctx, csr)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *http.Cursor) error); ok {
		r0 = rf(ctx, csr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
