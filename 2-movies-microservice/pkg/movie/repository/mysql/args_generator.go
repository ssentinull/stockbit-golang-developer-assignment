package mysql

import "github.com/ssentinull/stockbit-assignment/pkg/utils/http"

func genCreateSearchLogArgs(csr *http.Cursor) []interface{} {
	args := make([]interface{}, 0)
	args = append(args, csr.GetPage(), csr.GetSearchWord())

	return args
}
