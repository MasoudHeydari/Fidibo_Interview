package v1

import (
	"github.com/MasoudHeydari/Fidibo_Interview/contract"
	"github.com/MasoudHeydari/Fidibo_Interview/dto"
	"github.com/MasoudHeydari/Fidibo_Interview/interactor/book"
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
