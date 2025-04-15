package repository

import (
	"backend/internal/entity"
	"database/sql"
)

type UserRepository interface {
	FindByUsername(username string) (*entity.User, error)
	CreateUser(user *entity.User) error
	GetAllUsers() ([]*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(userID int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByUsername(username string) (*entity.User, error) {
	query := `
		SELECT u.id, u.username, u.password, u.name, d.id, d.department
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
		&user.Name,
		&user.Department.ID,
		&user.Department.Department,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user *entity.User) error {
	query := "INSERT INTO users (username, password, name, department_id) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, user.Username, user.Password, user.Name, user.Department.ID)
	return err
}

func (r *userRepository) GetAllUsers() ([]*entity.User, error) {
	query := `
		SELECT u.id, u.username, u.password, d.id, u.name, d.department
		FROM users u
		LEFT JOIN tb_department d ON u.department_id = d.id
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		user := &entity.User{Department: &entity.Department{}}
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Department.ID, &user.Name, &user.Department.Department)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) UpdateUser(user *entity.User) error {
	query := "UPDATE users SET password=$1, department_id=$2 WHERE username=$3"
	_, err := r.db.Exec(query, user.Password, user.Department.ID, user.Username)
	return err
}

func (r *userRepository) DeleteUser(userID int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.db.Exec(query, userID)
	return err
}
