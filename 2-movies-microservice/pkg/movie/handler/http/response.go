package http

import "github.com/ssentinull/stockbit-assignment/pkg/domain"

type movieResponse struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	IMDBID string `json:"imdb_id"`
	Poster string `json:"poster"`
}

func newMovieResponses(mvs []domain.Movie) []*movieResponse {
	mvResponses := make([]*movieResponse, 0)
	for _, mv := range mvs {
		mvResponses = append(mvResponses, &movieResponse{
			Title:  mv.Title,
			Year:   mv.Year,
			IMDBID: mv.IMDBID,
			Poster: mv.Poster,
		})
	}

	return mvResponses
}
