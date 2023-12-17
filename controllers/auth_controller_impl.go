package controllers

import (
	"midtrans-go/helper"
	"midtrans-go/model/web"
	"midtrans-go/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (service *AuthControllerImpl) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	authRequestDto := web.AuthRequestDto{
		Username: username,
		Password: password,
	}

	authService := service.AuthService.Authenticate(r.Context(), authRequestDto)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   authService,
	}

	helper.WriteToResponse(w, webResponse)
}

func (service *AuthControllerImpl) CreateJWT(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	authService := service.AuthService.CreateJWT(r.Context(), web.AuthRequestDto{})

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   authService,
	}

	helper.WriteToResponse(w, webResponse)
}
