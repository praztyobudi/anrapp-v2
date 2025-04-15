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
	query := `
		SELECT u.id, u.username, u.password, d.id, d.name, d.department
		FROM users u
		LEFT JOIN tb_department d ON u.department_id = d.id
		WHERE u.username = $1
	`
	row := r.db.QueryRow(query, username)

	user := &entity.User{Department: &entity.Department{}}
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Department.ID,
		&user.Department.Name,
		&user.Department.Department,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
