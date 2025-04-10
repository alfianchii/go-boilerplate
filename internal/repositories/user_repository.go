package repositories

import (
	"context"
	"errors"
	"go-boilerplate/internal/database"
	"go-boilerplate/internal/models"
)

type UserRepositoryInterface interface {
	FindByUsername(ctx context.Context, username string) (*models.User, error)
}

type UserRepository struct {
	DB *database.DB
}

func NewUserRepository(db *database.DB) UserRepositoryInterface {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `SELECT * FROM users WHERE username = $1`
	row := r.DB.Pool.QueryRow(ctx,query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.CreatedBy, &user.UpdatedAt, &user.UpdatedBy, &user.DeletedAt, &user.DeletedBy)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}