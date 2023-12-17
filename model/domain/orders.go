package domain

import "time"

type Orders struct {
	Id        int
	OrderId   int
	Amount    float64
	Status    string
	SnapUrl   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
