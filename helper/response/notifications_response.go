package response

import (
	"midtrans-go/model/domain"
	"midtrans-go/model/web"
)

func ToNotificationsResponse(notification domain.Notifications) web.NotificationsResponse {
	return web.NotificationsResponse{
		Id:               notification.Id,
		Status:           notification.Status,
		MidtransResponse: notification.MidtransResponse,
		CreatedAt:        notification.CreatedAt,
		UpdatedAt:        notification.UpdatedAt,
	}
}

func ToNotificationsResponses(order []domain.Notifications) []web.NotificationsResponse {
	var notificationResponses []web.NotificationsResponse
	for _, order := range order {
		notificationResponses = append(notificationResponses, ToNotificationsResponse(order))
	}
	return notificationResponses
}
