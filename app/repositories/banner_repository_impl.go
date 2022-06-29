package repositories

import (
	"go-app/app/entities"

	"gorm.io/gorm"
)

type bannerRepositoryImpl struct {
	db *gorm.DB
}

func NewBannerRepository(db *gorm.DB) *bannerRepositoryImpl {
	return &bannerRepositoryImpl{db}
}

func (r *bannerRepositoryImpl) FindAll() ([]entities.Banner, error) {
	var banners []entities.Banner
	err := r.db.Find(&banners).Error
	if err != nil {
		return banners, err
	}

	return banners, nil
}

func (r *bannerRepositoryImpl) FindById(bannerID int) (entities.Banner, error) {
	var banner entities.Banner
	err := r.db.First(&banner, bannerID).Error
	if err != nil {
		return banner, err
	}

	return banner, nil
}

func (r *bannerRepositoryImpl) Save(banner entities.Banner) (entities.Banner, error) {
	err := r.db.Create(&banner).Error
	if err != nil {
		return banner, err
	}

	return banner, nil
}

func (r *bannerRepositoryImpl) Update(bannerID int) (entities.Banner, error) {
	var model entities.Banner
	banner, err := r.FindById(bannerID)
	if err != nil {
		return banner, err
	}

	var updatedBanner bool
	if banner.IsShow {
		updatedBanner = false
	} else {
		updatedBanner = true
	}

	err = r.db.Model(&model).Where("id = ?", bannerID).Update("is_show", updatedBanner).Error
	if err != nil {
		return banner, err
	}

	return banner, nil
}

func (r *bannerRepositoryImpl) Delete(bannerID int) (entities.Banner, error) {
	var banner entities.Banner
	err := r.db.Delete(banner, bannerID).Error
	if err != nil {
		return banner, err
	}

	return banner, nil
}
