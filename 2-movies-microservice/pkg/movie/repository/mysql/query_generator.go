package mysql

func genCreateSearchLogQry() string {
	query := `INSER search_log SET page = ?, search_word = ?, created_at = ?`

	return query
}
