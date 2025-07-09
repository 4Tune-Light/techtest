package repository

import (
	"gitlab.com/rizkyimaduddin24/techtest/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(datas *[]entity.User) error
	GetById(data *entity.User, id string) error
	GetByEmail(data *entity.User, email string) error
	Create(data *entity.User) error
	Update(data *entity.User, id string) error
	Delete(data *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) GetAll(datas *[]entity.User) error {
	repo.db.Find(datas)
	return repo.db.Error
}

func (repo *userRepository) GetById(data *entity.User, id string) error {
	repo.db.First(data, "id = ?", id)
	return repo.db.Error
}

func (repo *userRepository) GetByEmail(data *entity.User, email string) error {
	repo.db.First(data, "email = ?", email)
	return repo.db.Error
}

func (repo *userRepository) Create(data *entity.User) error {
	repo.db.Create(data)
	return repo.db.Error
}

func (repo *userRepository) Update(data *entity.User, id string) error {
	repo.db.Model(data)
	repo.db.Where("id = ?", id)
	repo.db.Save(data)
	return repo.db.Error
}

func (repo *userRepository) Delete(data *entity.User) error {
	repo.db.Delete(data)
	return repo.db.Error
}
