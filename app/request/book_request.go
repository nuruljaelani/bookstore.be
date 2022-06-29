package request

type BookRequest struct {
	Title      string `json:"title" binding:"required"`
	Price      string `json:"price" binding:"required"`
	Author     string `json:"author" binding:"required"`
	Publisher  string `json:"publisher" binding:"required"`
	Stock      int    `json:"stock" binding:"numeric"`
	Thick      int    `json:"thick" binding:"numeric"`
	Desc       string `json:"desc" binding:"required"`
	Slug       string
	CategoryID int `json:"category" binding:"required"`
}
