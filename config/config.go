package config

import (
	"midtrans-go/helper"
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Port        string
	Environment string
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

	var midtransConfig Midtrans

	if os.Getenv("ENVIRONMENT") == "prod" {
		midtransConfig = Midtrans{
			MidtransId: os.Getenv("MIDTRANS_ID_PROD"),
			ClientKey:  os.Getenv("CLIENT_KEY_PROD"),
			ServerKey:  os.Getenv("SERVER_KEY_PROD"),
		}
	} else {
		midtransConfig = Midtrans{
			MidtransId: os.Getenv("MIDTRANS_ID_DEV"),
			ClientKey:  os.Getenv("CLIENT_KEY_DEV"),
			ServerKey:  os.Getenv("SERVER_KEY_DEV"),
		}
	}

	return &Config{
		App{
			Port:        os.Getenv("PORT"),
			Environment: os.Getenv("ENVIRONMENT"),
		},
		Database{
			Driver:   os.Getenv("DB_DRIVER"),
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
		},
		midtransConfig,
	}
}
