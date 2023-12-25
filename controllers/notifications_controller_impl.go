package controllers

import (
	"midtrans-go/helper"
	"midtrans-go/model/web"
	"midtrans-go/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type NotificationsImpl struct {
	NotificationsService service.NotificationsService
}

func NewNotificationsController(notificationsService service.NotificationsService) *NotificationsImpl {
	return &NotificationsImpl{
		NotificationsService: notificationsService,
	}
}

func (controller *NotificationsImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	notificationsRequestDto := web.NotificationsRequestDto{}
	helper.ReadFromRequestBody(r, &notificationsRequestDto)

	createOrders := controller.NotificationsService.Create(r.Context(), notificationsRequestDto)

	helper.WriteToResponse(w, createOrders)
}

func (controller *NotificationsImpl) GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	notificationId := p.ByName("notificationId")
	id, err := strconv.Atoi(notificationId)

	helper.PanicIfError(
		"Cannot catch notification id",
		err,
	)

	getByNotificationsId := controller.NotificationsService.GetById(r.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   getByNotificationsId,
	}

	helper.WriteToResponse(w, webResponse)
}

func (controller *NotificationsImpl) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	getNotifications := controller.NotificationsService.Get(r.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   getNotifications,
	}

	helper.WriteToResponse(w, webResponse)
}
