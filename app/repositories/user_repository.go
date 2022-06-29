package repositories

import "go-app/app/entities"

type UserRepository interface {
	FindAll() ([]entities.User, error)
	FindById(userID uint) (entities.User, error)
	Save(user entities.User) (entities.User, error)
}
