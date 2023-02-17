package main

import (
	v1 "fidibo_interview/delivery/http/v1"
	"github.com/labstack/echo/v4"
)

func main() {
	// init redis cache

	e := echo.New()
	e.GET("/search/book", v1.SearchBook())
	e.Logger.Fatal(e.Start(":8080"))
}
