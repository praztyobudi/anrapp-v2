package repository

import (
	"backend/internal/entity"
	"context"
	"database/sql"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	FindDepartmentByName(ctx context.Context, name string) (*entity.Department, error)
	CreateUser(ctx context.Context, user *entity.User) error
	GetAllUsers(ctx context.Context) ([]*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, userID int) error // Update signature di sini
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

// FindByUsername retrieves a user by username
func (r *userRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := `
		SELECT u.id, u.username, u.password, u.name, d.id, d.department
		FROM users u
		LEFT JOIN tb_department d ON u.department_id = d.id
		WHERE u.username = $1
	`
	row := r.db.QueryRowContext(ctx, query, username)

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
		if err == sql.ErrNoRows {
			return nil, nil // username not found
		}
		return nil, err // Other errors
	}
	return user, nil
}
func (r *userRepository) FindDepartmentByName(ctx context.Context, name string) (*entity.Department, error) {
	query := "SELECT id, department FROM tb_department WHERE department = $1"
	row := r.db.QueryRowContext(ctx, query, name)

	var dept entity.Department
	err := row.Scan(&dept.ID, &dept.Department)
	if err != nil {
		return nil, err
	}
	return &dept, nil
}

// CreateUser inserts a new user into the database
func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) error {
	query := "INSERT INTO users (username, password, name, department_id) VALUES ($1, $2, $3, $4)"
	_, err := r.db.ExecContext(ctx, query, user.Username, user.Password, user.Name, user.Department.ID)
	return err
}

// GetAllUsers retrieves all users from the database
func (r *userRepository) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	query := `
		SELECT u.id, u.username, u.password, u.name, d.id, d.department
		FROM users u
		LEFT JOIN tb_department d ON u.department_id = d.id
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		user := &entity.User{Department: &entity.Department{}}
		err := rows.Scan(
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
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUser updates user information in the database
func (r *userRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	query := "UPDATE users SET username = $1, password = $2, name = $3, department_id = $4 WHERE id = $5"
	_, err := r.db.ExecContext(ctx, query, user.Username, user.Password, user.Name, user.Department.ID, user.ID)
	return err
}

// DeleteUser deletes a user from the database
func (r *userRepository) DeleteUser(ctx context.Context, userID int) error { // Perbaikan di sini
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}
