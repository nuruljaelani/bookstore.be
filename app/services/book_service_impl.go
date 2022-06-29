package services

import (
	"go-app/app/entities"
	"go-app/app/repositories"
	"go-app/app/request"
	"strings"
)

type bookServiceImpl struct {
	repository repositories.BookRepository
}

func NewBookService(repository repositories.BookRepository) *bookServiceImpl {
	return &bookServiceImpl{repository}
}

func (s *bookServiceImpl) FindAll() ([]entities.Book, error) {
	books, err := s.repository.FindAll()
	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *bookServiceImpl) FindById(bookID int) (entities.Book, error) {
	book, err := s.repository.FindById(bookID)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *bookServiceImpl) Create(request request.BookRequest) (entities.Book, error) {
	book := entities.Book{}
	book.Title = request.Title
	book.Author = request.Author
	book.Price = request.Price
	book.Publisher = request.Publisher
	book.Stock = request.Stock
	book.Thick = request.Thick
	book.Desc = request.Desc
	strToLower := strings.ToLower(book.Title)
	strToSlug := strings.ReplaceAll(strToLower, " ", "-")
	book.Slug = strToSlug
	book.CategoryID = request.CategoryID

	newBook, err := s.repository.Save(book)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (s *bookServiceImpl) Update(bookID int, request request.BookRequest) (entities.Book, error) {
	book, err := s.repository.FindById(bookID)
	if err != nil {
		return book, err
	}

	book.Title = request.Title
	book.Author = request.Author
	book.Price = request.Price
	book.Publisher = request.Publisher
	book.Stock = request.Stock
	book.Thick = request.Thick
	book.Desc = request.Desc
	strToLower := strings.ToLower(request.Slug)
	strToSlug := strings.ReplaceAll(strToLower, " ", "-")
	book.Slug = strToSlug

	updatedBook, err := s.repository.Update(book)
	if err != nil {
		return updatedBook, err
	}

	return updatedBook, nil
}

func (s *bookServiceImpl) Delete(bookID int) (entities.Book, error) {
	deletedBook, err := s.repository.Delete(bookID)
	if err != nil {
		return deletedBook, err
	}

	return deletedBook, nil
}

func (s *bookServiceImpl) Upload(bookID int, fileLocation string, mime string) (entities.BookImage, error) {
	bookImage := entities.BookImage{}
	bookImage.BookID = uint(bookID)
	bookImage.FileName = fileLocation
	bookImage.Mime = mime

	imageUploaded, err := s.repository.Upload(bookImage)
	if err != nil {
		return imageUploaded, err
	}

	return imageUploaded, nil
}
