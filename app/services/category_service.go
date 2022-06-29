package services

import (
	"go-app/app/entities"
	"go-app/app/request"
)

type CategoryService interface {
	FindAll() ([]entities.Category, error)
	FindById(categoryID int) (entities.Category, error)
	Create(request request.CategoryRequest) (entities.Category, error)
	Update(categoryID int, request request.CategoryRequest) (entities.Category, error)
	Delete(categoryID int) (entities.Category, error)
}
