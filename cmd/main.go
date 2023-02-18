package main

import (
	"github.com/MasoudHeydari/Fidibo_Interview/adapter/store"
	"github.com/MasoudHeydari/Fidibo_Interview/cache"
	"github.com/MasoudHeydari/Fidibo_Interview/config"
	v1 "github.com/MasoudHeydari/Fidibo_Interview/delivery/http/v1"
	"github.com/MasoudHeydari/Fidibo_Interview/validator"
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

	// initialize Echo
	e := echo.New()
	e.POST("/register", v1.CreateUser(mysqlStore, validator.ValidateCreateUser))
	e.POST("/login", v1.LoginUser(mysqlStore, validator.ValidateLoginUser))
	e.POST("/search/book", v1.SearchBook(bookCache), v1.AuthMiddleware)
	e.Logger.Fatal(e.Start(config.GetHttpAddress()))
}
