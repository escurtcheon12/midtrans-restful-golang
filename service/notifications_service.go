package service

import (
	"context"
	"midtrans-go/model/web"
)

type NotificationsService interface {
	Create(ctx context.Context, request web.NotificationsRequestDto) web.NotificationsResponse
	GetById(ctx context.Context, notificationId int) web.NotificationsResponse
	Get(ctx context.Context) []web.NotificationsResponse
}
