package mysql

func genCreateGetMoviesLogQry() string {
	query := `INSER search_log SET page = ?, search_word = ?, created_at = ?`

	return query
}

func genCreateGetMovieByTitleLogQry() string {
	query := `INSER search_log SET title = ?, created_at = ?`

	return query
}
