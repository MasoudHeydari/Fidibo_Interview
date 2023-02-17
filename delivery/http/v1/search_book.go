package v1

import (
	"fidibo_interview/dto"
	"fidibo_interview/interactor"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SearchBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		q := c.QueryParam("keyword")
		req := dto.SearchBookRequest{
			Keyword: q,
		}

		resp, err := interactor.New().SearchBook(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}
