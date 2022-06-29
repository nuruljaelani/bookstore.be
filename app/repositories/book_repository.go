package repositories

import (
	"go-app/app/entities"
)

type BookRepository interface {
	FindAll() ([]entities.Book, error)
	FindById(bookID int) (entities.Book, error)
	Save(book entities.Book) (entities.Book, error)
	Update(book entities.Book) (entities.Book, error)
	Delete(bookID int) (entities.Book, error)
	Upload(bookImage entities.BookImage) (entities.BookImage, error)
}
