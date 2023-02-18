package main

import (
	"fidibo_interview/adapter/store"
	"fidibo_interview/cache"
	"fidibo_interview/config"
	v1 "fidibo_interview/delivery/http/v1"
	"fidibo_interview/validator"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	// connect to database and auto migrate
	mysqlStore := store.New(config.GetStoreDsn())

	// init redis cache
	cacheDur, err := config.GetRedisCacheDuration()
	if err != nil {
		log.Fatal("failed to extract the cache duration - err: ", err)
	}
	bookCache := cache.NewRedisCache(config.GetRedisDsn(), 1, cacheDur)
	fmt.Println("cache started")

	// initialize Echo
	e := echo.New()
	e.POST("/register", v1.CreateUser(mysqlStore, validator.ValidateCreateUser))
	e.POST("/login", v1.LoginUser(mysqlStore, validator.ValidateLoginUser))
	e.POST("/search/book", v1.SearchBook(bookCache), v1.AuthMiddleware)
	e.Logger.Fatal(e.Start(config.GetHttpAddress()))
}
