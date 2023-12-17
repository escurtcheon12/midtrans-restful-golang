package service

import (
	"context"
	"midtrans-go/model/web"
)

type AuthService interface {
	Authenticate(ctx context.Context, user web.AuthRequestDto) []web.AuthResponse
	CreateJWT(ctx context.Context, user web.AuthRequestDto) string
}
