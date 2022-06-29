package repositories

import "go-app/app/entities"

type BannerRepository interface {
	FindAll() ([]entities.Banner, error)
	FindById(bannerID int) (entities.Banner, error)
	Save(banner entities.Banner) (entities.Banner, error)
	Update(bannerID int) (entities.Banner, error)
	Delete(bannerID int) (entities.Banner, error)
}
