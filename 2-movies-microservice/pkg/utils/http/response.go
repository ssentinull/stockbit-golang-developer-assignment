package http

import "github.com/ssentinull/stockbit-assignment/pkg/utils"

type CursorResponse struct {
	Page       string      `json:"page"`
	SearchWord string      `json:"search_word"`
	Title      string      `json:"title"`
	Payload    interface{} `json:"payload"`
}

func NewCursorResponse(csr *Cursor, p interface{}) *CursorResponse {
	return &CursorResponse{
		Page:       utils.IntToString(csr.GetPage()),
		SearchWord: csr.GetSearchWord(),
		Title:      csr.GetTitle(),
		Payload:    p,
	}
}
