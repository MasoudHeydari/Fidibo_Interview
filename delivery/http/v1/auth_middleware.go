package v1

import (
	"errors"
	"fmt"
	"github.com/MasoudHeydari/Fidibo_Interview/auth"
	"github.com/MasoudHeydari/Fidibo_Interview/config"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

var (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header not provided")
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		authorizationFields := strings.Fields(authorizationHeader)
		if len(authorizationFields) < 2 {
			err := errors.New("invalid authorization header format")
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		authorizationType := strings.ToLower(authorizationFields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		accessToken := authorizationFields[1]
		tokenMaker, err := auth.NewJWtMaker(config.GetSecretKey())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		_, err = tokenMaker.VerifyTokenPayload(accessToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		return next(c)
	}
}
