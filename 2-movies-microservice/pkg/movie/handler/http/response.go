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

type movieDetailsResponse struct {
	Title               string                         `json:"title"`
	ReleaseYear         string                         `json:"release_year"`
	AgeRating           string                         `json:"age_rating"`
	ReleaseDate         string                         `json:"release_date"`
	Runtime             string                         `json:"tuntime"`
	Genre               []string                       `json:"genres"`
	Directors           []string                       `json:"directors"`
	Writers             []string                       `json:"writers"`
	Actors              []string                       `json:"actors"`
	Plot                string                         `json:"plot"`
	Languages           []string                       `json:"languages"`
	Countries           []string                       `json:"countries"`
	Awards              string                         `json:"awards"`
	Poster              string                         `json:"poster"`
	Ratings             []*movieDetailsRatingsResponse `json:"ratings"`
	Metascore           string                         `json:"metascore"`
	IMDBRating          string                         `json:"imdb_rating"`
	IMDBVotes           string                         `json:"imdb_votes"`
	IMDBID              string                         `json:"imdb_id"`
	Type                string                         `json:"type"`
	DVDReleaseDate      string                         `json:"dvd_release_date"`
	BoxOffice           string                         `json:"box_office"`
	ProductionCompanies []string                       `json:"production_companies"`
	Website             string                         `json:"website"`
	ErrorMessage        string                         `json:"error_message"`
}

type movieDetailsRatingsResponse struct {
	Source string `json:"source"`
	Value  string `json:"value"`
}

func newMovieDetailsResponse(md domain.MovieDetails) *movieDetailsResponse {
	ratingsResponses := make([]*movieDetailsRatingsResponse, 0)
	for _, rating := range md.Ratings {
		ratingsResponses = append(ratingsResponses, &movieDetailsRatingsResponse{
			Source: rating.Source,
			Value:  rating.Value,
		})
	}

	return &movieDetailsResponse{
		Title:               md.Title,
		ReleaseYear:         md.ReleaseYear,
		AgeRating:           md.AgeRating,
		ReleaseDate:         md.ReleaseDate,
		Runtime:             md.Runtime,
		Genre:               md.Genre,
		Directors:           md.Directors,
		Writers:             md.Writers,
		Actors:              md.Actors,
		Plot:                md.Plot,
		Languages:           md.Languages,
		Countries:           md.Countries,
		Awards:              md.Awards,
		Poster:              md.Poster,
		Ratings:             ratingsResponses,
		Metascore:           md.Metascore,
		IMDBRating:          md.IMDBRating,
		IMDBVotes:           md.IMDBVotes,
		IMDBID:              md.IMDBID,
		Type:                md.Type,
		DVDReleaseDate:      md.DVDReleaseDate,
		BoxOffice:           md.BoxOffice,
		ProductionCompanies: md.ProductionCompanies,
		Website:             md.Website,
		ErrorMessage:        md.ErrorMessage,
	}
}
