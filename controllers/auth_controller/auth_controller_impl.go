package authcontroller

import (
	"midtrans-go/helper"
	"midtrans-go/model/web"
	authweb "midtrans-go/model/web/auth_web"
	service "midtrans-go/service/auth_service"
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

	authRequestDto := authweb.AuthRequestDto{
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
	authService := service.AuthService.CreateJWT(r.Context(), authweb.AuthRequestDto{})

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   authService,
	}

	helper.WriteToResponse(w, webResponse)
}
