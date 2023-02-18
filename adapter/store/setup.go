package store

import (
	"fidibo_interview/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type MySQLStore struct {
	db *gorm.DB
}

func New(dsn string) MySQLStore {
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	q := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;",
		config.GetDbName(),
	)
	database.Exec(q)

	if err != nil {
		panic("Failed to connect to database!")
	}

	if aErr := database.AutoMigrate(&User{}); aErr != nil {
		panic("Failed to auto migrate database!")
	}

	log.Println("connected to database successfully")
	return MySQLStore{db: database}
}
