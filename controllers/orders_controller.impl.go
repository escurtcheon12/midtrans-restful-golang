package controllers

import (
	"midtrans-go/helper"
	"midtrans-go/model/web"
	"midtrans-go/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type OrdersControllerImpl struct {
	OrdersService service.OrdersService
}

func NewOrdersController(ordersService service.OrdersService) *OrdersControllerImpl {
	return &OrdersControllerImpl{
		OrdersService: ordersService,
	}
}

func (controller *OrdersControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ordersRequestDto := web.OrdersRequestDto{}
	helper.ReadFromRequestBody(r, &ordersRequestDto)

	createOrders := controller.OrdersService.Create(r.Context(), ordersRequestDto)

	helper.WriteToResponse(w, createOrders)
}

func (controller *OrdersControllerImpl) GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	orderId := p.ByName("orderId")
	id, err := strconv.Atoi(orderId)

	helper.PanicIfError(
		"Cannot catch order id",
		err,
	)

	getByOrdersId := controller.OrdersService.GetById(r.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   getByOrdersId,
	}

	helper.WriteToResponse(w, webResponse)
}

func (controller *OrdersControllerImpl) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	getOrders := controller.OrdersService.Get(r.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   getOrders,
	}

	helper.WriteToResponse(w, webResponse)
}
