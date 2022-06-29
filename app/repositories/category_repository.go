package repositories

import "go-app/app/entities"

type CategoryRepository interface {
	FindAll() ([]entities.Category, error)
	FindById(categoryID int) (entities.Category, error)
	Save(category entities.Category) (entities.Category, error)
	Update(category entities.Category) (entities.Category, error)
	Delete(categoryID int) (entities.Category, error)
}
