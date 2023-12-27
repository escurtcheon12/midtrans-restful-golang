package domain

import (
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

type Orders struct {
	Id               int
	Status           string
	MidtransResponse coreapi.ChargeResponse
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
