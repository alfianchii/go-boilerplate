package models

import "time"

type Role struct {
	ID int `json:"role_id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *string `json:"deleted_by"`
}