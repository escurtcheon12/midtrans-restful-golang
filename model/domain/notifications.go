package domain

import "time"

type Notifications struct {
	Id               int
	Status           string
	MidtransResponse string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
