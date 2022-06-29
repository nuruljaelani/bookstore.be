package services

import (
	"go-app/app/entities"
	"go-app/app/repositories"
)

type bannerServiceImpl struct {
	repository repositories.BannerRepository
}

func NewBannerService(repository repositories.BannerRepository) *bannerServiceImpl {
	return &bannerServiceImpl{repository}
}

func (s *bannerServiceImpl) FindAll() ([]entities.Banner, error) {
	banners, err := s.repository.FindAll()
	if err != nil {
		return banners, err
	}

	return banners, nil
}

func (s *bannerServiceImpl) FindById(bannerID int) (entities.Banner, error) {
	banner, err := s.repository.FindById(bannerID)
	if err != nil {
		return banner, err
	}

	return banner, nil
}

func (s *bannerServiceImpl) Create(fileLoction, mime string) (entities.Banner, error) {
	banner := entities.Banner{}
	banner.FileName = fileLoction
	banner.Mime = mime

	newBanner, err := s.repository.Save(banner)
	if err != nil {
		return newBanner, err
	}

	return newBanner, nil
}

func (s *bannerServiceImpl) Delete(bannerId int) (entities.Banner, error) {
	deletedBanner, err := s.repository.Delete(bannerId)
	if err != nil {
		return deletedBanner, err
	}

	return deletedBanner, nil
}

func (s *bannerServiceImpl) UpdateShowBanner(bannerId int) (entities.Banner, error) {
	updatedBanner, err := s.repository.Update(bannerId)
	if err != nil {
		return updatedBanner, err
	}

	return updatedBanner, nil
}
