package pg

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
	"webproj/internal/domain"
)

type UsersPgRepo struct {
	db *DB
}

func NewUserPgRepo(db *DB) *UsersPgRepo {
	return &UsersPgRepo{db: db}
}

func (r *UsersPgRepo) Create(ctx context.Context, user *domain.User) error {
	query := `
        INSERT INTO users (first_name,last_name,email,password,created_at)
        VALUES ($1,$2,$3,$4,$5)
    `
	_, err := r.db.Exec(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		time.Now(),
	)
	if err != nil {
		return fmt.Errorf("User with this email already exists")
	}
	return nil
}

func (r *UsersPgRepo) Delete(ctx context.Context, userId uuid.UUID) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(ctx, query, userId)
	if err != nil {
		fmt.Printf("failed to delete user %v \n", err)
	}
	return nil
}

func (r *UsersPgRepo) Get(ctx context.Context, emailOrId interface{}) (*domain.User, error) {
	var user domain.User
	var query string
	switch v := emailOrId.(type) {
	case string:
		err := uuid.Validate(v)
		if err != nil {
			query = `SELECT * FROM users WHERE email=$1`
			err := r.db.QueryRow(ctx, query, v).Scan(
				&user.Id,
				&user.FirstName,
				&user.LastName,
				&user.Email,
				&user.Password,
				&user.CreatedAt)
			if err != nil {
				return nil, fmt.Errorf("failed to find user by email")
			}
			return &user, nil
		} else {
			query = `SELECT * FROM users WHERE id=$1`
			err := r.db.QueryRow(ctx, query, v).Scan(
				&user.Id,
				&user.FirstName,
				&user.LastName,
				&user.Email,
				&user.Password,
				&user.CreatedAt)
			if err != nil {
				return nil, fmt.Errorf("failed to find user by id")
			}
			return &user, nil
		}
	case uuid.UUID:
		query = `SELECT * FROM users WHERE id=$1`
		err := r.db.QueryRow(ctx, query, v).Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to find user by id")
		}
		return &user, nil
	default:
		return nil, fmt.Errorf("Invalid input data")
	}
}
