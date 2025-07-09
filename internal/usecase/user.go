package usecase

import (
	"errors"

	"gitlab.com/rizkyimaduddin24/techtest/internal/entity"
	"gitlab.com/rizkyimaduddin24/techtest/internal/repository"
)

type UserUsecase interface {
	GetAll() ([]entity.User, error)
	GetByID(id string) (entity.User, error)
	Update(id string, name string, email string) error
	Delete(id string) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (uc *userUsecase) GetAll() ([]entity.User, error) {
	// Validate Email
	users := &[]entity.User{}
	err := uc.repo.GetAll(users)
	if len(*users) < 1 || err != nil {
		return nil, errors.New("user not found")
	}

	return *users, nil
}

func (uc *userUsecase) GetByID(id string) (entity.User, error) {
	// Validate ID
	user := &entity.User{}
	err := uc.repo.GetById(user, id)
	if user.ID == 0 || err != nil {
		return *user, errors.New("user not found")
	}

	return *user, nil
}

func (uc *userUsecase) Update(id string, name string, email string) error {
	// Validate ID
	user := &entity.User{}
	err := uc.repo.GetById(user, id)
	if user.ID == 0 || err != nil {
		return errors.New("user not found")
	}

	// Update Data
	user.Name = name
	user.Email = email
	if err := uc.repo.Update(user, id); err != nil {
		return errors.New("failed to update user")
	}

	return nil
}

func (uc *userUsecase) Delete(id string) error {
	// Validate ID
	user := &entity.User{}
	err := uc.repo.GetById(user, id)
	if user.ID == 0 || err != nil {
		return errors.New("user not found")
	}

	// Delete Data
	if err := uc.repo.Delete(user); err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}
