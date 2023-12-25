package repositories

import (
	"context"
	"database/sql"
	"midtrans-go/model/domain"
)

type OrdersRepository interface {
	Create(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders
	Get(ctx context.Context, tx *sql.Tx) []domain.Orders
	GetById(ctx context.Context, tx *sql.Tx, orderId int) (domain.Orders, error)
}
