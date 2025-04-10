package repositories

import (
	"context"
	"errors"
	"go-boilerplate/internal/database"
	"go-boilerplate/internal/models"
)

type UserRepositoryInterface interface {
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	FindByUsernameWithRoles(ctx context.Context, username string) (*models.User, error)
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

func (r *UserRepository) FindByUsernameWithRoles(ctx context.Context, username string) (*models.User, error) {
	query := `
		SELECT
			u.*,
			r.*
		FROM users u
		LEFT JOIN user_roles ur ON u.user_id = ur.user_id
		LEFT JOIN roles r ON ur.role_id = r.role_id
		WHERE u.username = $1
	`

	rows, err := r.DB.Pool.Query(ctx, query, username)
	if err != nil {
		return nil, errors.New("failed to query user with roles")
	}
	defer rows.Close()

	var user *models.User
	var roles []models.Role

	for rows.Next() {
		var role models.Role
		if user == nil {
			user = &models.User{}
			err := rows.Scan(
				&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.CreatedBy, &user.UpdatedAt, &user.UpdatedBy, &user.DeletedAt, &user.DeletedBy,
				&role.ID, &role.Name, &role.CreatedAt, &role.CreatedBy, &role.UpdatedAt, &role.UpdatedBy, &role.DeletedAt, &role.DeletedBy,
			)
			if err != nil {
				return nil, errors.New("failed to scan user and role")
			}
		} else {
			err := rows.Scan(
				nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
				&role.ID, &role.Name, &role.CreatedAt, &role.CreatedBy, &role.UpdatedAt, &role.UpdatedBy, &role.DeletedAt, &role.DeletedBy,
			)
			if err != nil {
				return nil, errors.New("failed to scan role data")
			}
		}

		roles = append(roles, role)
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	user.Roles = roles
	return user, nil
}