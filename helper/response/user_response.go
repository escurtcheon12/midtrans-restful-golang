package response

import (
	"midtrans-go/model/domain"
	"midtrans-go/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(user []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range user {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
