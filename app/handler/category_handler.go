package handler

import (
	"go-app/app/helper"
	"go-app/app/request"
	"go-app/app/response"
	"go-app/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	service services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) *categoryHandler {
	return &categoryHandler{service}
}

func (h *categoryHandler) GetCategories(c *gin.Context) {
	categories, err := h.service.FindAll()
	if err != nil {
		response := helper.APIResponse("Error to get categories", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of categories", http.StatusOK, "success", response.ToCategoryResponses(categories))
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) GetCategoryById(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Param("id"))

	category, err := h.service.FindById(categoryId)
	if err != nil {
		response := helper.APIResponse("Failed to get detail category", http.StatusNotFound, "error", err.Error())
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Get detail category", http.StatusOK, "success", response.ToCategoryResponse(category))
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var input request.CategoryRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorsMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create category not success", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	category, err := h.service.Create(input)
	if err != nil {
		response := helper.APIResponse("Create category not success", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Create category successfull", http.StatusOK, "success", response.ToCategoryResponse(category))
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	var req request.CategoryRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorsMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to update category", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	categoryId, _ := strconv.Atoi(c.Param("id"))
	updatedCategory, err := h.service.Update(categoryId, req)
	if err != nil {
		response := helper.APIResponse("Failed to update category", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update category successfull", http.StatusOK, "success", response.ToCategoryResponse(updatedCategory))
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Param("id"))

	_, err := h.service.FindById(categoryId)
	if err != nil {
		response := helper.APIResponse("Failed to delete category", http.StatusNotFound, "error", err.Error())
		c.JSON(http.StatusNotFound, response)
		return
	}
	_, err = h.service.Delete(categoryId)
	if err != nil {
		response := helper.APIResponse("Failed to delete category", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete category successfull", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
