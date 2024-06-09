package main

import (
	"go-gin-productManagerPro/controllers"
	"go-gin-productManagerPro/models"
	"go-gin-productManagerPro/repositories"
	"go-gin-productManagerPro/services"

	"github.com/gin-gonic/gin"
)

func main() {
	products := []models.Product{
		{ID: 1, Name: "Product1", Price: 1000, Description: "Description1", SoldOut: false},
		{ID: 2, Name: "Product2", Price: 2000, Description: "Description2", SoldOut: true},
		{ID: 3, Name: "Product3", Price: 3000, Description: "Description3", SoldOut: false},
	}

	productRepository := repositories.NewProductMemoryRepository(products)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)


	r := gin.Default()
	r.GET("/products", productController.FindAll) 
	r.GET("/products/:id", productController.FindById)
	r.POST("/products", productController.Create)
	r.PUT("/products/:id", productController.Update)
	r.Run("localhost:8080")
}
