package model

type User struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Role     []string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,minLength=6,maxLength=24"`
	Password string `json:"password" validate:"required,minLength=8,maxLength=24"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,minLength=6,maxLength=24"`
	Password string `json:"password" validate:"required,minLength=8,maxLength=24"`
}
