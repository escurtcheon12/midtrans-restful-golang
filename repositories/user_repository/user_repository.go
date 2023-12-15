package userrepository

import (
	"context"
	"database/sql"
	"midtrans-go/model/domain"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, tx *sql.Tx, user domain.User) []domain.User
}
