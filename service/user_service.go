package service

import (
	"context"
	"midtrans-go/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserRequestDto) web.UserResponse
	GetById(ctx context.Context, userId int) web.UserResponse
	Get(ctx context.Context) []web.UserResponse
}
