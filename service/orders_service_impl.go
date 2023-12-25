package service

import (
	"context"
	"database/sql"
	"midtrans-go/helper"
	"midtrans-go/helper/response"
	"midtrans-go/model/domain"
	"midtrans-go/model/web"
	"midtrans-go/repositories"

	"github.com/go-playground/validator/v10"
)

type OrdersServiceImpl struct {
	OrdersRepository repositories.OrdersRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewOrdersService(ordersRepository repositories.OrdersRepository, DB *sql.DB, validator *validator.Validate) *OrdersServiceImpl {
	return &OrdersServiceImpl{
		OrdersRepository: ordersRepository,
		DB:               DB,
		Validate:         validator,
	}
}

func (service *OrdersServiceImpl) Create(ctx context.Context, request web.OrdersRequestDto) web.OrdersResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(
		"Error filder field validation",
		err,
	)

	tx, err := service.DB.Begin()
	helper.PanicIfError(
		"Cannot start db begin",
		err,
	)
	defer helper.CommitOrRollback(tx)

	user := domain.Orders{
		Status:           request.Status,
		MidtransResponse: request.MidtransResponse,
	}

	userRepository := service.OrdersRepository.Create(ctx, tx, user)

	return response.ToOrdersResponse(userRepository)
}

func (service *OrdersServiceImpl) GetById(ctx context.Context, orderId int) web.OrdersResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError("Cannot start DB Begin", err)

	getOrdersById, err := service.OrdersRepository.GetById(ctx, tx, orderId)
	if err != nil {
		helper.PanicIfError("Get by id is error", err)
	}

	return response.ToOrdersResponse(getOrdersById)
}

func (service *OrdersServiceImpl) Get(ctx context.Context) []web.OrdersResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError("Cannot start DB Begin", err)

	getOrders := service.OrdersRepository.Get(ctx, tx)

	return response.ToOrdersResponses(getOrders)
}
