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

type NotificationsServiceImpl struct {
	NotificationsRepository repositories.NotificationsRepository
	DB                      *sql.DB
	Validate                *validator.Validate
}

func NewNotificationsService(notificationsRepository repositories.NotificationsRepository, DB *sql.DB, validator *validator.Validate) *NotificationsServiceImpl {
	return &NotificationsServiceImpl{
		NotificationsRepository: notificationsRepository,
		DB:                      DB,
		Validate:                validator,
	}
}

func (service *NotificationsServiceImpl) Create(ctx context.Context, request web.NotificationsRequestDto) web.NotificationsResponse {
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

	notification := domain.Notifications{
		Status:           request.Status,
		MidtransResponse: web.NotificationCallbackRequestDto{},
	}

	userRepository := service.NotificationsRepository.Create(ctx, tx, notification)

	return response.ToNotificationsResponse(userRepository)
}

func (service *NotificationsServiceImpl) GetById(ctx context.Context, notificationId int) web.NotificationsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError("Cannot start DB Begin", err)

	getNotificationsById, err := service.NotificationsRepository.GetById(ctx, tx, notificationId)
	if err != nil {
		helper.PanicIfError("Get by id is error", err)
	}

	return response.ToNotificationsResponse(getNotificationsById)
}

func (service *NotificationsServiceImpl) Get(ctx context.Context) []web.NotificationsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError("Cannot start DB Begin", err)

	getNotifications := service.NotificationsRepository.Get(ctx, tx)

	return response.ToNotificationsResponses(getNotifications)
}
