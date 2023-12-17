package repositories

import (
	"context"
	"database/sql"
	"midtrans-go/model/domain"
)

type OrdersRepositoryImpl struct {
}

func NewOrdersRepository() *OrdersRepositoryImpl {
	return &OrdersRepositoryImpl{}
}

func Create(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders {
	return orders
}
