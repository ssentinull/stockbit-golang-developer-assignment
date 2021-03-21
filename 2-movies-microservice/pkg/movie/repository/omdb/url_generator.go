package omdb

import (
	"fmt"

	httpUtils "github.com/ssentinull/stockbit-assignment/pkg/utils/http"
)

func genReadMoviesURL(baseURL, key string, csr *httpUtils.Cursor) string {
	return fmt.Sprintf("%sapikey=%s&page=%v&s=%s", baseURL, key, csr.GetPage(), csr.GetSearchWord())
}

func genReadMovieByTitle(baseURL, key string, csr *httpUtils.Cursor) string {
	return fmt.Sprintf("%sapikey=%s&t=%s", baseURL, key, csr.GetTitle())
}
