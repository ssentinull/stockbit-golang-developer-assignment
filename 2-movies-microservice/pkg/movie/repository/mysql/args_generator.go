package mysql

import "github.com/ssentinull/stockbit-assignment/pkg/utils/http"

func genCreateGetMoviesLogArgs(csr *http.Cursor) []interface{} {
	args := make([]interface{}, 0)
	args = append(args, csr.GetPage(), csr.GetSearchWord())

	return args
}

func genCreateGetMovieByTitleLogArg(csr *http.Cursor) interface{} {
	return csr.GetTitle()
}
