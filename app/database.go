package app

import (
	"database/sql"
	"midtrans-go/config"
	"midtrans-go/helper"
	"time"
)

func NewDatabase() *sql.DB {
	DB_DRIVER := config.NewConfig().Database.Driver
	DB_USER := config.NewConfig().Database.User
	DB_PASSWORD := config.NewConfig().Database.Password
	DB_HOST := config.NewConfig().Database.Host
	DB_PORT := config.NewConfig().Database.Port
	DB_NAME := config.NewConfig().Database.Name

	db, err := sql.Open(DB_DRIVER, DB_USER+":"+DB_PASSWORD+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)
	helper.PanicIfError("Cannot connect to database", err)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
