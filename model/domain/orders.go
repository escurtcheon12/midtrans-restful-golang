package domain

import "time"

type Orders struct {
	Id               int
	Status           string
	MidtransResponse string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
