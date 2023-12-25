package repositories

import (
	"context"
	"database/sql"
	"midtrans-go/model/domain"
)

type NotificationsRepository interface {
	Create(ctx context.Context, tx *sql.Tx, notifications domain.Notifications) domain.Notifications
	Get(ctx context.Context, tx *sql.Tx) []domain.Notifications
	GetById(ctx context.Context, tx *sql.Tx, notificationId int) (domain.Notifications, error)
}
