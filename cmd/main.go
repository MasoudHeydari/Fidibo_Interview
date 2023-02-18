package main

import (
	"fidibo_interview/adapter/store"
	"fidibo_interview/cache"
	"fidibo_interview/config"
	v1 "fidibo_interview/delivery/http/v1"
	"fidibo_interview/validator"
	"github.com/labstack/echo/v4"
	"time"
)

func main() {
	// connect to database and auto migrate
	mysqlStore := store.New(config.GetMySqlDSN())

	// init redis cache
	bookCache := cache.NewRedisCache("localhost:6379", 1, time.Minute*10)

	e := echo.New()
	e.POST("/register", v1.CreateUser(mysqlStore, validator.ValidateCreateUser))
	e.POST("/login", v1.LoginUser(mysqlStore, validator.ValidateLoginUser))
	e.POST("/search/book", v1.SearchBook(bookCache), v1.AuthMiddleware)
	e.Logger.Fatal(e.Start(":8080"))
}
