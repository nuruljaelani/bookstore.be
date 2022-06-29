package request

type BookImageRequest struct {
	BookID int `form:"book_id" binding:"required"`
	Mime   string
}
