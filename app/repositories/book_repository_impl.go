package repositories

import (
	"go-app/app/entities"

	"gorm.io/gorm"
)

type bookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepositoryImpl {
	return &bookRepositoryImpl{db}
}

func (repository *bookRepositoryImpl) FindAll() ([]entities.Book, error) {
	var books []entities.Book
	err := repository.db.Preload("BookImages").Preload("Category").Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}

func (repository *bookRepositoryImpl) FindById(bookID int) (entities.Book, error) {
	var book entities.Book
	err := repository.db.Joins("Category").First(&book, bookID).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (repository *bookRepositoryImpl) Save(book entities.Book) (entities.Book, error) {
	err := repository.db.Create(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (repository *bookRepositoryImpl) Update(book entities.Book) (entities.Book, error) {
	err := repository.db.Save(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (repository *bookRepositoryImpl) Delete(bookID int) (entities.Book, error) {
	var book entities.Book
	err := repository.db.Delete(book, bookID).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (repository *bookRepositoryImpl) Upload(image entities.BookImage) (entities.BookImage, error) {
	err := repository.db.Create(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}
