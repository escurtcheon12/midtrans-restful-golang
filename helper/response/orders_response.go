package response

import (
	"midtrans-go/model/domain"
	"midtrans-go/model/web"
)

func ToOrdersResponse(order domain.Orders) web.OrdersResponse {
	return web.OrdersResponse{
		Id:               order.Id,
		Status:           order.Status,
		MidtransResponse: order.MidtransResponse,
		CreatedAt:        order.CreatedAt,
		UpdatedAt:        order.UpdatedAt,
	}
}

func ToOrdersResponses(order []domain.Orders) []web.OrdersResponse {
	var orderResponses []web.OrdersResponse
	for _, order := range order {
		orderResponses = append(orderResponses, ToOrdersResponse(order))
	}
	return orderResponses
}
