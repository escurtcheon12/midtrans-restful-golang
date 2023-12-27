package domain

import (
	"midtrans-go/model/web"
	"time"
)

type Notifications struct {
	Id               int
	Status           string
	MidtransResponse web.NotificationCallbackRequestDto
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
