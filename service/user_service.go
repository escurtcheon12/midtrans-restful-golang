package service

import (
	"database/sql"
	"midtrans-go/repositories"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserSevice(userRepository repositories.UserRepository, DB *sql.DB, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}
