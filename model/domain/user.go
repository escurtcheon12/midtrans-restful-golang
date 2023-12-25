package domain

import "time"

type User struct {
	Id        int
	Username  string
	Email     string
	Phone     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
