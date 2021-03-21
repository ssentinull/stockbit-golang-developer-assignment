package http

import (
	"github.com/labstack/echo/v4"
	"github.com/ssentinull/stockbit-assignment/pkg/utils"
)

type Cursor struct {
	page       int
	searchWord string
}

func (c *Cursor) GetPage() int {
	return c.page
}

func (c *Cursor) GetSearchWord() string {
	return c.searchWord
}

func NewCursor(c echo.Context) (*Cursor, error) {
	var err error
	searchWord := c.QueryParam("searchword")
	if searchWord == "" {
		err = ErrBadQueryParams

		return nil, err
	}

	page := utils.StringToInt(c.QueryParam("pagination"))
	if page == 0 {
		page = 1
	}

	csr := &Cursor{
		page:       page,
		searchWord: searchWord,
	}

	return csr, nil
}
