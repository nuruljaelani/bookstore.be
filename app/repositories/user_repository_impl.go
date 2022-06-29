package repositories

import (
	"go-app/app/entities"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db}
}

func (repository *userRepositoryImpl) FindAll() ([]entities.User, error) {
	var users []entities.User
	err := repository.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (repository *userRepositoryImpl) FindById(userID uint) (entities.User, error) {
	var user entities.User
	err := repository.db.Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
