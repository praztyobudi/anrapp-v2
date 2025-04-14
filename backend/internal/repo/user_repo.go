package repository

import (
	"backend/internal/entity"
	"database/sql"
)

type UserRepository interface {
	FindByUsername(username string) (*entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByUsername(username string) (*entity.User, error) {
	row := r.db.QueryRow("SELECT id, username, password FROM users WHERE username=$1", username)
	user := &entity.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
