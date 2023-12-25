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

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repositories.UserRepository, DB *sql.DB, validator *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validator,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserRequestDto) web.UserResponse {
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

	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: request.Password,
	}

	userRepository := service.UserRepository.Create(ctx, tx, user)

	return response.ToUserResponse(userRepository)
}

func (service *UserServiceImpl) GetById(ctx context.Context, userId int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError("Cannot start DB Begin", err)

	getUserById, err := service.UserRepository.GetById(ctx, tx, userId)
	if err != nil {
		helper.PanicIfError("Get by id is error", err)
	}

	return response.ToUserResponse(getUserById)
}

func (service *UserServiceImpl) Get(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError("Cannot start DB Begin", err)

	getNotifications := service.UserRepository.Get(ctx, tx)

	return response.ToUserResponses(getNotifications)
}
