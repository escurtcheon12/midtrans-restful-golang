package controllers

import (
	"midtrans-go/helper"
	"midtrans-go/model/web"
	"midtrans-go/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userRequestDto := web.UserRequestDto{}
	helper.ReadFromRequestBody(r, &userRequestDto)

	createUser := controller.UserService.Create(r.Context(), userRequestDto)

	helper.WriteToResponse(w, createUser)
}

func (controller *UserControllerImpl) GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("userId")
	id, err := strconv.Atoi(userId)

	helper.PanicIfError(
		"Cannot catch user id",
		err,
	)

	getByUserId := controller.UserService.GetById(r.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   getByUserId,
	}

	helper.WriteToResponse(w, webResponse)
}

func (controller *UserControllerImpl) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	getUser := controller.UserService.Get(r.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   getUser,
	}

	helper.WriteToResponse(w, webResponse)
}
