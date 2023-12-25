package web

type UserRequestDto struct {
	Username string `validate:"required" json:"username"`
	Email    string `validate:"required" json:"email"`
	Phone    string `validate:"required" json:"phone"`
	Password string `validate:"required" json:"password"`
}
