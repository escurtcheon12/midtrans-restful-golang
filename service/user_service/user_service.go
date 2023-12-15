package userservice

import (
	"database/sql"
	userrepository "midtrans-go/repositories/user_repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository userrepository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserSevice(userRepository userrepository.UserRepository, DB *sql.DB, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

// func authenticate(ctx context.Context, request user.User) {

// }
