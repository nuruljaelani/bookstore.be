package handler

import (
	"fmt"
	"go-app/app/helper"
	"go-app/app/response"
	"go-app/app/services"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bannerHandler struct {
	service services.BannerService
}

func NewBannerHandler(service services.BannerService) *bannerHandler {
	return &bannerHandler{service}
}

func (h *bannerHandler) GetBanners(c *gin.Context) {
	banners, err := h.service.FindAll()
	if err != nil {
		res := helper.APIResponse("Failed to get banners", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.APIResponse("List of banners", http.StatusOK, "success", response.ToBannerResponses(banners))
	c.JSON(http.StatusOK, res)
}

func (h *bannerHandler) GetBannerById(c *gin.Context) {
	bannerId, _ := strconv.Atoi(c.Param("id"))
	banner, err := h.service.FindById(bannerId)
	if err != nil {
		res := helper.APIResponse("Failed to get banners", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.APIResponse("List of banners", http.StatusOK, "error", response.ToBannerResponse(banner))
	c.JSON(http.StatusOK, res)
}

func (h *bannerHandler) CreateBanner(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", err.Error())

		c.JSON(http.StatusBadRequest, response)
		return
	}
	files := form.File["banner[]"]

	for _, file := range files {
		path := fmt.Sprintf("images/banners/%s", file.Filename)
		mime := filepath.Ext(path)

		err := c.SaveUploadedFile(file, path)
		if err != nil {
			response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", err.Error())

			c.JSON(http.StatusBadRequest, response)
			return
		}

		_, err = h.service.Create(path, mime)
		if err != nil {
			response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", err.Error())

			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Image has been uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *bannerHandler) DeleteBanner(c *gin.Context) {
	bannerId, _ := strconv.Atoi(c.Param("id"))

	banner, err := h.service.FindById(bannerId)
	if err != nil {
		response := helper.APIResponse("Failed to delete banner", http.StatusNotFound, "error", err.Error())
		c.JSON(http.StatusNotFound, response)
		return
	}

	_, err = h.service.Delete(bannerId)
	if err != nil {
		response := helper.APIResponse("Failed to delete banner", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = os.Remove(banner.FileName)
	if err != nil {
		response := helper.APIResponse("Failed to delete banner", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"deleted": true}
	response := helper.APIResponse("Delete banner successfull", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *bannerHandler) UpdateBanner(c *gin.Context) {
	bannerId, _ := strconv.Atoi(c.Param("id"))

	_, err := h.service.UpdateShowBanner(bannerId)
	if err != nil {
		response := helper.APIResponse("Failed to update banner", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"updated": true}
	response := helper.APIResponse("Update banner successfull", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
