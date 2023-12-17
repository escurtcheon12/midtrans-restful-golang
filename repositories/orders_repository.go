package repositories

import (
	"context"
	"database/sql"
	"midtrans-go/model/domain"
)

type OrdersRepository interface {
	Create(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders
}
