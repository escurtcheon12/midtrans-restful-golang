package web

import "time"

type NotificationsResponse struct {
	Id               int       `json:"id"`
	Status           string    `json:"status"`
	MidtransResponse string    `json:"midtrans_response"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
