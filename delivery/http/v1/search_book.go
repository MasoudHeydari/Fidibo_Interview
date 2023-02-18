package v1

import (
	"fidibo_interview/contract"
	"fidibo_interview/dto"
	"fidibo_interview/interactor/book"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SearchBook(redisCache contract.BookCache) echo.HandlerFunc {
	return func(c echo.Context) error {
		q := c.QueryParam("keyword")
		req := dto.SearchBookRequest{
			Keyword: q,
		}

		resp, err := book.New(redisCache).SearchBook(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}
