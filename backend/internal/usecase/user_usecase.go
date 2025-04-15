package usecase

import (
	"backend/internal/entity"
	repository "backend/internal/repo"
	"errors"
	"strings"
)

type UserUsecase interface {
	Login(username, password string) (*entity.User, error)
	Register(user *entity.User) error
	GetUsers() ([]*entity.User, error)
	Update(user *entity.User) error
	Delete(userID int) error
}

type userUsecaseImpl struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecaseImpl{repo}
}

func (u *userUsecaseImpl) Login(username, password string) (*entity.User, error) {
	user, err := u.repo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if password != "" && user.Password != password {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (u *userUsecaseImpl) Register(user *entity.User) error {
	if strings.TrimSpace(user.Username) == "" ||
		strings.TrimSpace(user.Password) == "" ||
		strings.TrimSpace(user.Name) == "" ||
		strings.TrimSpace(user.Department.Department) == "" {
		return errors.New("username, password, and name are required")
	}
	existing, _ := u.repo.FindByUsername(user.Username)
	if existing != nil {
		return errors.New("username already exists")
	}
	return u.repo.CreateUser(user)
}

func (u *userUsecaseImpl) GetUsers() ([]*entity.User, error) {
	return u.repo.GetAllUsers()
}

func (u *userUsecaseImpl) Update(user *entity.User) error {
	if strings.TrimSpace(user.Password) == "" || user.Department.ID == 0 {
		return errors.New("password and department_id are required")
	}
	return u.repo.UpdateUser(user)
}

func (u *userUsecaseImpl) Delete(userID int) error {
	return u.repo.DeleteUser(userID)
}
