package response

import (
	"midtrans-go/model/domain"
	"midtrans-go/model/web"
)

func ToAuthResponse(user domain.User) web.AuthResponse {
	return web.AuthResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToAuthResponses(user []domain.User) []web.AuthResponse {
	var authResponses []web.AuthResponse
	for _, user := range user {
		authResponses = append(authResponses, ToAuthResponse(user))
	}
	return authResponses
}
