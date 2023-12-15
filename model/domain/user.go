package domain

import "time"

type User struct {
	Id         int
	Username   string
	Email      string
	Phone      string
	Password   string
	Created_at time.Time
	Updated_at time.Time
}
