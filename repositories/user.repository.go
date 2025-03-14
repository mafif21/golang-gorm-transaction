package repositories

import (
	"golang-transaction-experiment/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(id int) (*model.User, error)
	GetAll() ([]model.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r UserRepositoryImpl) GetUserById(id int) (*model.User, error) {
	var user *model.User

	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return user, nil
}

func (r UserRepositoryImpl) GetAll() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
