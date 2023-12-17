package main

import (
	"midtrans-go/app"
	"midtrans-go/config"
	"midtrans-go/controllers"
	"midtrans-go/middleware"
	"midtrans-go/repositories"
	"midtrans-go/service"
	"net/http"

	"github.com/midtrans/midtrans-go/coreapi"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func main() {
	db := app.NewDatabase()
	validate := validator.New()
	userRepository := repositories.NewUserRepository()
	authService := service.NewAuthService(userRepository, db, validate)
	authController := controllers.NewAuthController(authService)

	midtransService := service.NewMidtransService()
	midtransController := controllers.NewMidtransController(coreapi.Client{}, midtransService)

	router := app.NewRouter(authController, midtransController)

	server := http.Server{
		Addr:    "localhost:" + config.NewConfig().App.Port,
		Handler: middleware.NewMiddleware(router),
	}

	logrus.Info("Running at port " + config.NewConfig().App.Port)
	server.ListenAndServe()

}
