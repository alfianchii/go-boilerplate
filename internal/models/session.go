package models

import "time"

type Session struct {
	RowID string `json:"rowid"`
	UserID int `json:"user_id"`
	Token string `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	IPAddress string `json:"ip_address"`
	IsBlacklisted bool `json:"is_blacklisted"`
}