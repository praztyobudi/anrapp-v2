package usecase

import (
	"backend/internal/dto"
	"backend/internal/entity"
	repository "backend/internal/repo"
	"context"
	"database/sql"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Login(ctx context.Context, username, password string) (*entity.User, error)
	Register(ctx context.Context, req *dto.RegisterRequest) error
	GetUsers(ctx context.Context) ([]*entity.User, error)
	Update(ctx context.Context, req *dto.UpdateUserRequest) error
	Delete(ctx context.Context, userID int) error
}

type userUsecaseImpl struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecaseImpl{repo}
}

func (u *userUsecaseImpl) Login(ctx context.Context, username, password string) (*entity.User, error) {
	user, err := u.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (u *userUsecaseImpl) Register(ctx context.Context, req *dto.RegisterRequest) error {
	if strings.TrimSpace(req.Username) == "" ||
		strings.TrimSpace(req.Password) == "" ||
		strings.TrimSpace(req.Name) == "" ||
		strings.TrimSpace(req.DepartmentName) == "" {
		return errors.New("username, password, name, and department name are required")
	}

	// Cek username
	existing, err := u.repo.FindByUsername(ctx, req.Username)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if existing != nil {
		return errors.New("username already exists")
	}

	// Cari department berdasarkan nama
	dept, err := u.repo.FindDepartmentByName(ctx, req.DepartmentName)
	if err != nil {
		return errors.New("department not found")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := &entity.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Name:     req.Name,
		Department: &entity.Department{
			ID: dept.ID,
		},
	}

	return u.repo.CreateUser(ctx, user)
}

func (u *userUsecaseImpl) GetUsers(ctx context.Context) ([]*entity.User, error) {
	return u.repo.GetAllUsers(ctx)
}

func (u *userUsecaseImpl) Update(ctx context.Context, req *dto.UpdateUserRequest) error {
	if strings.TrimSpace(req.Username) == "" ||
		strings.TrimSpace(req.Password) == "" ||
		strings.TrimSpace(req.Name) == "" ||
		strings.TrimSpace(req.DepartmentName) == "" {
		return errors.New("all fields are required")
	}

	// cari ID department berdasarkan nama
	dept, err := u.repo.FindDepartmentByName(ctx, req.DepartmentName)
	if err != nil {
		return errors.New("invalid department name")
	}

	// mapping ke entity
	user := &entity.User{
		ID:       req.ID,
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Department: &entity.Department{
			ID: dept.ID,
		},
	}

	return u.repo.UpdateUser(ctx, user)
}

func (u *userUsecaseImpl) Delete(ctx context.Context, userID int) error {
	return u.repo.DeleteUser(ctx, userID)
}
