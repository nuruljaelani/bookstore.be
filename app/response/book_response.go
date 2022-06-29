package response

import "go-app/app/entities"

type bookResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Price     string `json:"price"`
	Stock     int    `json:"stock"`
	Thick     int    `json:"thick"`
	Desc      string `json:"desc"`
	Slug      string `json:"slug"`
	Category  string `json:"category"`
	Image     string `json:"image"`
}

func ToBookResponse(book entities.Book) bookResponse {
	bookResponse := bookResponse{}
	bookResponse.ID = int(book.ID)
	bookResponse.Title = book.Title
	bookResponse.Author = book.Author
	bookResponse.Publisher = book.Publisher
	bookResponse.Price = book.Price
	bookResponse.Stock = book.Stock
	bookResponse.Thick = book.Thick
	bookResponse.Desc = book.Desc
	bookResponse.Slug = book.Slug
	bookResponse.Category = book.Category.Name
	bookResponse.Image = ""

	if len(book.BookImages) > 0 {
		bookResponse.Image = book.BookImages[0].FileName
	}

	return bookResponse
}

func ToBookResponses(books []entities.Book) []bookResponse {
	booksResponse := []bookResponse{}
	for _, book := range books {
		bookResponse := ToBookResponse(book)
		booksResponse = append(booksResponse, bookResponse)
	}

	return booksResponse
}
