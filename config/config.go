package config

import (
	"midtrans-go/helper"
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Port string
}

type Database struct {
	Driver   string
	Host     string
	User     string
	Password string
	Port     string
	Name     string
}

type Midtrans struct {
	MidtransId string
	ClientKey  string
	ServerKey  string
}

type Config struct {
	App
	Database
	Midtrans
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	helper.PanicIfError("Cannot load .env", err)

	return &Config{
		App{
			Port: os.Getenv("PORT"),
		},
		Database{
			Driver:   os.Getenv("DB_DRIVER"),
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
		},
		Midtrans{
			MidtransId: os.Getenv("MIDTRANS_ID"),
			ClientKey:  os.Getenv("CLIENT_KEY"),
			ServerKey:  os.Getenv("SERVER_KEY"),
		},
	}
}
