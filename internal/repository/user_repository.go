package repository

import (
	"database/sql"
	"fmt"

	"github.com/Oliver1ck/testApi/internal/models"
)




type UserRepository interface {
  CreateUser(user *models.User) (*models.User, error)
  GetUserByUsername(username string) (*models.User, error)
}

type userRepository struct {
  db *sql.DB
}


func NewUserRepository(db *sql.DB) UserRepository {
  return &userRepository{db: db}
}


func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
  query := `
		INSERT INTO users (username, email, password_hash, created_at)
		VALUES ($1, $2, $3, NOW())
		RETURNING id, created_at
	`
	
	err := r.db.QueryRow(
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.CreatedAt)
	
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
  return user, nil
}


func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
  return nil, nil
}