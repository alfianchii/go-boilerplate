package models

import "time"

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type User struct {
	ID int `json:"user_id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *int `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *int `json:"deleted_by"`
	Roles []Role `json:"roles"`
}