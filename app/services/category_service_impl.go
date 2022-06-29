package services

import (
	"go-app/app/entities"
	"go-app/app/repositories"
	"go-app/app/request"
)

type categoryServiceImpl struct {
	repository repositories.CategoryRepository
}

func NewCategoryService(repository repositories.CategoryRepository) *categoryServiceImpl {
	return &categoryServiceImpl{repository}
}

func (s *categoryServiceImpl) FindAll() ([]entities.Category, error) {
	categories, err := s.repository.FindAll()
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (s *categoryServiceImpl) FindById(categoryID int) (entities.Category, error) {
	category, err := s.repository.FindById(categoryID)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *categoryServiceImpl) Create(request request.CategoryRequest) (entities.Category, error) {
	category := entities.Category{}
	category.Name = request.Name

	newCategory, err := s.repository.Save(category)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *categoryServiceImpl) Update(categoriID int, request request.CategoryRequest) (entities.Category, error) {
	category, err := s.repository.FindById(categoriID)
	if err != nil {
		return category, err
	}

	category.Name = request.Name
	updatedCategory, err := s.repository.Update(category)
	if err != nil {
		return updatedCategory, err
	}

	return updatedCategory, nil
}

func (s *categoryServiceImpl) Delete(categoryID int) (entities.Category, error) {
	category, err := s.repository.Delete(categoryID)
	if err != nil {
		return category, err
	}

	return category, nil
}
