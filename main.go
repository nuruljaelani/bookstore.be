package main

import (
	"go-app/app/handler"
	"go-app/app/repositories"
	"go-app/app/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=root dbname=bookstore port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	bookRepository := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	bannerRepository := repositories.NewBannerRepository(db)
	bannerService := services.NewBannerService(bannerRepository)
	bannerHandler := handler.NewBannerHandler(bannerService)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	// Route Category
	api.GET("/categories", categoryHandler.GetCategories)
	api.GET("/categories/:id", categoryHandler.GetCategoryById)
	api.POST("/categories/create", categoryHandler.CreateCategory)
	api.PUT("/categories/update/:id", categoryHandler.UpdateCategory)
	api.DELETE("/categories/delete/:id", categoryHandler.DeleteCategory)

	// Route Book
	api.GET("/books", bookHandler.GetBooks)
	api.GET("/books/:id", bookHandler.GetBookById)
	api.POST("/books/create", bookHandler.CreateBook)
	api.PUT("/books/update/:id", bookHandler.UpdateBook)
	api.DELETE("/books/delete/:id", bookHandler.DeleteBook)
	api.POST("/books/:id/upload", bookHandler.UploadImage)

	// Route Banner
	api.GET("/banners", bannerHandler.GetBanners)
	api.POST("/banners/create", bannerHandler.CreateBanner)
	api.DELETE("/banners/delete/:id", bannerHandler.DeleteBanner)
	api.PUT("/banners/:id/update", bannerHandler.UpdateBanner)

	router.Run()

	// fmt.Println("Connected Successfully")
	// db.AutoMigrate(&entities.Category{})
	// db.AutoMigrate(&entities.User{})
	// db.AutoMigrate(&entities.Book{})
	// db.AutoMigrate(&entities.BookImage{})
	// db.AutoMigrate(&entities.Banner{})

}
