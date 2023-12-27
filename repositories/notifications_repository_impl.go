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

type NotificationsRepositoryImpl struct {
}

func NewNotificationsRepository() *NotificationsRepositoryImpl {
	return &NotificationsRepositoryImpl{}
}

func (repository *NotificationsRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, notifications domain.Notifications) domain.Notifications {
	notifResponse, err := json.Marshal(notifications.MidtransResponse)
	helper.PanicIfError("Error marshaling data notif", err)

	SQL := "insert into notifications(status,midtrans_response) values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, notifications.MidtransResponse.TransactionStatus, string(notifResponse))
	helper.PanicIfError("Error to create data", err)

	id, err := result.LastInsertId()
	helper.PanicIfError("Error to catch last id", err)

	notifications.CreatedAt = time.Now()
	notifications.UpdatedAt = time.Now()
	notifications.Id = int(id)
	return notifications
}

func (repository *NotificationsRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, notificationId int) (domain.Notifications, error) {
	SQL := "select id, status, midtrans_response, created_at, updated_at from notifications where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, notificationId)
	helper.PanicIfError("Error to get data", err)
	defer rows.Close()

	notification := domain.Notifications{}
	if rows.Next() {
		var createdAtStr, updatedAtStr string

		err := rows.Scan(&notification.Id, &notification.Status, &notification.MidtransResponse, &createdAtStr, &updatedAtStr)
		if err != nil {
			helper.PanicIfError("Error to convert data", err)
		}

		notification.CreatedAt = utils.ParseTimestamp(createdAtStr)
		notification.UpdatedAt = utils.ParseTimestamp(updatedAtStr)

		return notification, nil
	} else {
		return notification, errors.New("category is not found")
	}
}

func (repository *NotificationsRepositoryImpl) Get(ctx context.Context, tx *sql.Tx) []domain.Notifications {
	SQL := "SELECT id, status, midtrans_response, created_at, updated_at FROM notifications"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError("Error to get data", err)
	defer rows.Close()

	var notifications []domain.Notifications
	for rows.Next() {
		notification := domain.Notifications{}
		var createdAtStr, updatedAtStr string

		err := rows.Scan(&notification.Id, &notification.Status, &notification.MidtransResponse, &createdAtStr, &updatedAtStr)

		if err != nil {
			helper.PanicIfError("Error to convert data", err)
		}

		notification.CreatedAt = utils.ParseTimestamp(createdAtStr)
		notification.UpdatedAt = utils.ParseTimestamp(updatedAtStr)

		notifications = append(notifications, notification)
	}
	return notifications
}
