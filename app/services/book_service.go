package services

import (
	"go-app/app/entities"
	"go-app/app/request"
)

type BookService interface {
	FindAll() ([]entities.Book, error)
	FindById(bookID int) (entities.Book, error)
	Create(request request.BookRequest) (entities.Book, error)
	Update(bookID int, request request.BookRequest) (entities.Book, error)
	Delete(bookID int) (entities.Book, error)
	Upload(bookID int, fileLocation string, mime string) (entities.BookImage, error)
}
