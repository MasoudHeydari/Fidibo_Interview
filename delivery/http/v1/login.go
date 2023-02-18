package v1

import (
	"fidibo_interview/adapter/store"
	"fidibo_interview/contract"
	"fidibo_interview/dto"
	"fidibo_interview/interactor/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func LoginUser(store store.MySQLStore, validator contract.ValidateLoginUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req = dto.LoginUserRequest{}
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := validator(req); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		resp, err := user.New(store).Login(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}
