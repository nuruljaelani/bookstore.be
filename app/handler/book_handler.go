package handler

import (
	"fmt"
	"go-app/app/helper"
	"go-app/app/request"
	"go-app/app/response"
	"go-app/app/services"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	service services.BookService
}

func NewBookHandler(service services.BookService) *bookHandler {
	return &bookHandler{service}
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.service.FindAll()
	if err != nil {
		response := helper.APIResponse("Error to get Books", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of Books", http.StatusOK, "success", response.ToBookResponses(books))
	c.JSON(http.StatusOK, response)
}

func (h *bookHandler) GetBookById(c *gin.Context) {
	bookId, _ := strconv.Atoi(c.Param("id"))
	book, err := h.service.FindById(bookId)
	if err != nil {
		response := helper.APIResponse("Error to get Book", http.StatusNotFound, "error", err.Error())
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Detail of Book", http.StatusOK, "success", response.ToBookResponse(book))
	c.JSON(http.StatusOK, response)
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var input request.BookRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorsMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create book is not success", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newBook, err := h.service.Create(input)
	if err != nil {
		response := helper.APIResponse("Create book is not success", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Create category successfull", http.StatusOK, "success", response.ToBookResponse(newBook))
	c.JSON(http.StatusOK, response)
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	var input request.BookRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorsMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to update category", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	bookId, _ := strconv.Atoi(c.Param("id"))
	updatedBook, err := h.service.Update(bookId, input)
	if err != nil {
		response := helper.APIResponse("Failed to update category", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update book successfull", http.StatusOK, "success", response.ToBookResponse(updatedBook))
	c.JSON(http.StatusOK, response)
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	bookId, _ := strconv.Atoi(c.Param("id"))

	_, err := h.service.FindById(bookId)
	if err != nil {
		response := helper.APIResponse("Failed to delete book", http.StatusNotFound, "error", err.Error())
		c.JSON(http.StatusNotFound, response)
		return
	}

	_, err = h.service.Delete(bookId)
	if err != nil {
		response := helper.APIResponse("Failed to delete book", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete book successfull", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *bookHandler) UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%s", file.Filename)
	bookID := c.Param("id")
	id, _ := strconv.Atoi(bookID)
	mime := filepath.Ext(path)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.Upload(id, path, mime)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Image has been uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
