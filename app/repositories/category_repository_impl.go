package repositories

import (
	"go-app/app/entities"

	"gorm.io/gorm"
)

type categoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepositoryImpl {
	return &categoryRepositoryImpl{db}
}

func (repository *categoryRepositoryImpl) Save(category entities.Category) (entities.Category, error) {
	err := repository.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (repository *categoryRepositoryImpl) FindAll() ([]entities.Category, error) {
	var categories []entities.Category
	err := repository.db.Find(&categories).Error
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (repository *categoryRepositoryImpl) FindById(categoryID int) (entities.Category, error) {
	var category entities.Category
	err := repository.db.First(&category, categoryID).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (repository *categoryRepositoryImpl) Update(category entities.Category) (entities.Category, error) {
	err := repository.db.Save(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (repository *categoryRepositoryImpl) Delete(categoryID int) (entities.Category, error) {
	var category entities.Category
	err := repository.db.Delete(category, categoryID).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
