package repositories

import (
	"context"
	"errors"
	"go-boilerplate/internal/database"
	"go-boilerplate/internal/models"
)

type SessionRepositoryInterface interface {
	StoreSession(ctx context.Context, session models.Session) error
	IsTokenBlacklisted(ctx context.Context, token string) (bool, error)
}

type SessionRepository struct {
	DB *database.DB
}

func NewSessionRepository(db *database.DB) SessionRepositoryInterface {
	return &SessionRepository{DB: db}
}

func (r *SessionRepository) StoreSession(ctx context.Context, session models.Session) error {
	query := `INSERT INTO sessions (rowid, user_id, token, expires_at, created_at, ip_address, is_blacklisted) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.DB.Pool.Exec(ctx, query,
		session.RowID,
		session.UserID,
		session.Token,
		session.ExpiresAt,
		session.CreatedAt,
		session.IPAddress,
		session.IsBlacklisted,
	)

	if err != nil {
		return errors.New("failed to store session")
	}
	return nil
}

func (r *SessionRepository) IsTokenBlacklisted(ctx context.Context, token string) (bool, error) {
	query := `SELECT is_blacklisted FROM sessions WHERE token = $1`
	row := r.DB.Pool.QueryRow(ctx, query, token)
	var isBlacklisted bool
	err := row.Scan(&isBlacklisted)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return false, nil
		}
		return false, errors.New("failed to check if token is blacklisted")
	}

	return isBlacklisted, nil
}