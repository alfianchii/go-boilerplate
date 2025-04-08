package models

import "time"

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type User struct {
	ID string `json:"user_id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *string `json:"deleted_by"`
}