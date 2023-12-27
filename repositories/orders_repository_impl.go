package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"midtrans-go/helper"
	"midtrans-go/model/domain"
	"midtrans-go/utils"
	"time"
)

type OrdersRepositoryImpl struct {
}

func NewOrdersRepository() *OrdersRepositoryImpl {
	return &OrdersRepositoryImpl{}
}

func (repository *OrdersRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders {
	ordersResponse, err := json.Marshal(orders.MidtransResponse)
	helper.PanicIfError("Error marshaling data orders", err)

	SQL := "insert into orders(status,midtrans_response) values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, orders.Status, string(ordersResponse))
	helper.PanicIfError("Error to create data", err)

	id, err := result.LastInsertId()
	helper.PanicIfError("Error to catch last id", err)

	orders.CreatedAt = time.Now()
	orders.UpdatedAt = time.Now()
	orders.Id = int(id)
	return orders
}

func (repository *OrdersRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, orderId int) (domain.Orders, error) {
	SQL := "select id, status, midtrans_response, created_at, updated_at from orders where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, orderId)
	helper.PanicIfError("Error to get data", err)
	defer rows.Close()

	order := domain.Orders{}
	if rows.Next() {
		var createdAtStr, updatedAtStr string

		err := rows.Scan(&order.Id, &order.Status, &order.MidtransResponse, &createdAtStr, &updatedAtStr)
		if err != nil {
			helper.PanicIfError("Error to convert data", err)
		}

		order.CreatedAt = utils.ParseTimestamp(createdAtStr)
		order.UpdatedAt = utils.ParseTimestamp(updatedAtStr)

		return order, nil
	} else {
		return order, errors.New("category is not found")
	}
}

func (repository *OrdersRepositoryImpl) Get(ctx context.Context, tx *sql.Tx) []domain.Orders {
	SQL := "SELECT id, status, midtrans_response, created_at, updated_at FROM orders"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError("Error to get data", err)
	defer rows.Close()

	var orders []domain.Orders
	for rows.Next() {
		order := domain.Orders{}
		var createdAtStr, updatedAtStr string

		err := rows.Scan(&order.Id, &order.Status, &order.MidtransResponse, &createdAtStr, &updatedAtStr)

		if err != nil {
			helper.PanicIfError("Error to convert data", err)
		}

		order.CreatedAt = utils.ParseTimestamp(createdAtStr)
		order.UpdatedAt = utils.ParseTimestamp(updatedAtStr)

		orders = append(orders, order)
	}
	return orders
}
