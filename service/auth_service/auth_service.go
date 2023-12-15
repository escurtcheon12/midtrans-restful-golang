package authservice

import (
	"context"
	authweb "midtrans-go/model/web/auth_web"
)

type AuthService interface {
	Authenticate(ctx context.Context, user authweb.AuthRequestDto) []authweb.AuthResponse
	CreateJWT(ctx context.Context, user authweb.AuthRequestDto) string
}
