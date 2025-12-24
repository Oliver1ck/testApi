package repository

import (
	"database/sql"

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
  
  return user, nil
}


func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
  return nil, nil
}