package helper

import (
	"midtrans-go/model/domain"
	authweb "midtrans-go/model/web/auth_web"
)

func ToAuthResponse(user domain.User) authweb.AuthResponse {
	return authweb.AuthResponse{
		Id:         user.Id,
		Username:   user.Username,
		Email:      user.Email,
		Phone:      user.Phone,
		Password:   user.Password,
		Created_at: user.Created_at,
		Updated_at: user.Updated_at,
	}
}

func ToAuthResponses(user []domain.User) []authweb.AuthResponse {
	var authResponses []authweb.AuthResponse
	for _, user := range user {
		authResponses = append(authResponses, ToAuthResponse(user))
	}
	return authResponses
}
