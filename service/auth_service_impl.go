package service

import (
	"context"
	"database/sql"
	"midtrans-go/helper"
	"midtrans-go/repositories"

	"time"

	"midtrans-go/model/domain"
	"midtrans-go/model/web"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type AuthServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthService(userRepository repositories.UserRepository, DB *sql.DB, validator *validator.Validate) *AuthServiceImpl {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validator,
	}
}

func (service *AuthServiceImpl) Authenticate(ctx context.Context, request web.AuthRequestDto) []web.AuthResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError("Error filder field validation",
		err,
	)

	tx, err := service.DB.Begin()
	helper.PanicIfError("Cannot start db begin",
		err,
	)

	user := domain.User{
		Username: request.Username,
		Password: request.Password,
	}

	userRepository := service.UserRepository.GetByUsername(ctx, tx, user)

	var authResponses []web.AuthResponse

	for _, u := range userRepository {
		token := service.CreateJWT(ctx, request)

		authResponses = append(authResponses, web.AuthResponse{
			Id:       u.Id,
			Username: u.Username,
			Email:    u.Email,
			Phone:    u.Phone,
			Password: u.Password,
			Token:    token,
		})
	}

	return authResponses
}

func (service *AuthServiceImpl) CreateJWT(ctx context.Context, request web.AuthRequestDto) string {
	token, err := generateToken()

	if err != nil {
		helper.PanicIfError("Generate token failed", err)
	}

	return token
}

func generateToken() (string, error) {
	var SECRET = []byte("secret-token")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "token"
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
		helper.PanicIfError("Token create err", err)
	}

	return tokenStr, nil
}
