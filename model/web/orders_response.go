package web

import (
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

type OrdersResponse struct {
	Id               int                    `json:"id"`
	Status           string                 `json:"status"`
	MidtransResponse coreapi.ChargeResponse `json:"midtrans_response"`
	CreatedAt        time.Time              `json:"created_at"`
	UpdatedAt        time.Time              `json:"updated_at"`
}
