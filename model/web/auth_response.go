package web

import "time"

type AuthResponse struct {
	Id         int
	Username   string
	Email      string
	Phone      string
	Password   string
	Created_at time.Time
	Updated_at time.Time
	Token      string
}
