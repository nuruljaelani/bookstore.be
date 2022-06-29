package response

import "go-app/app/entities"

type categoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToCategoryResponse(category entities.Category) categoryResponse {
	categoryResponse := categoryResponse{}
	categoryResponse.ID = category.ID
	categoryResponse.Name = category.Name

	return categoryResponse
}

func ToCategoryResponses(categories []entities.Category) []categoryResponse {

	categoriesResponse := []categoryResponse{}
	for _, category := range categories {
		categoryResponse := ToCategoryResponse(category)
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return categoriesResponse
}
