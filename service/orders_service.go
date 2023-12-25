package service

import (
	"context"
	"midtrans-go/model/web"
)

type OrdersService interface {
	Create(ctx context.Context, request web.OrdersRequestDto) web.OrdersResponse
	GetById(ctx context.Context, orderId int) web.OrdersResponse
	Get(ctx context.Context) []web.OrdersResponse
}
