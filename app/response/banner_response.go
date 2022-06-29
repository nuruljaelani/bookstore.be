package response

import "go-app/app/entities"

type bannerResponse struct {
	ID       int
	FileName string `json:"banner"`
	IsShow   bool   `json:"is_show"`
}

func ToBannerResponse(banner entities.Banner) bannerResponse {
	bannerRes := bannerResponse{}
	bannerRes.ID = banner.ID
	bannerRes.FileName = banner.FileName
	bannerRes.IsShow = banner.IsShow

	return bannerRes
}

func ToBannerResponses(banners []entities.Banner) []bannerResponse {
	bannerResponses := []bannerResponse{}
	for _, banner := range banners {
		bannerResponse := ToBannerResponse(banner)
		bannerResponses = append(bannerResponses, bannerResponse)
	}

	return bannerResponses
}
