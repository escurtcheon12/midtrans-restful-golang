package repositories

import (
	"context"
	"database/sql"
	"midtrans-go/model/domain"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, tx *sql.Tx, user domain.User) []domain.User
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Get(ctx context.Context, tx *sql.Tx) []domain.User
	GetById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
}
