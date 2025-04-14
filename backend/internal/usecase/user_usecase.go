package usecase

import (
	"backend/internal/entity"
	repository "backend/internal/repo"
	"errors"
)

type UserUsecase interface {
	Login(username, password string) (*entity.User, error)
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

	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
