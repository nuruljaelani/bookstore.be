package services

import "go-app/app/entities"

type BannerService interface {
	FindAll() ([]entities.Banner, error)
	FindById(bannerID int) (entities.Banner, error)
	Create(fileLocation, mime string) (entities.Banner, error)
	Delete(bannerID int) (entities.Banner, error)
	UpdateShowBanner(bannerID int) (entities.Banner, error)
}
