package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Roles     []string  `json:"roles"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,minLength=6,maxLength=24"`
	Password string `json:"password" validate:"required,minLength=8,maxLength=32"`
}

type RegisterRequest struct {
	Username        string `json:"username" validate:"required,minLength=6,maxLength=24"`
	Password        string `json:"password" validate:"required,minLength=8,maxLength=32"`
	ConfirmPassword string `json:"confirm_password" validate:"required,minLength=8,maxLength=32"`
}
