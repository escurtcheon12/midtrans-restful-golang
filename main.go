package main

import (
	"github.com/midtrans/midtrans-go/snap"
	"midtrans-go/app"
	"midtrans-go/config"
	authcontroller "midtrans-go/controllers/auth_controller"
	midtranscontroller "midtrans-go/controllers/midtrans_controller"
	"midtrans-go/middleware"
	userrepository "midtrans-go/repositories/user_repository"
	authservice "midtrans-go/service/auth_service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func main() {
	db := app.NewDatabase()
	validate := validator.New()
	userRepository := userrepository.NewUserRepository()
	authService := authservice.NewAuthService(userRepository, db, validate)
	authController := authcontroller.NewAuthController(authService)

	var snapClient snap.Client
	midtransController := midtranscontroller.NewMidtransController(snapClient)

	router := app.NewRouter(authController, midtransController)

	server := http.Server{
		Addr:    "localhost:" + config.NewConfig().App.Port,
		Handler: middleware.NewMiddleware(router),
	}

	// app.NewMidtransClient()

	logrus.Info("Running at port " + config.NewConfig().App.Port)
	server.ListenAndServe()

}
