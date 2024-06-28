package main

import (
	"go-gin-productManagerPro/controllers"
	"go-gin-productManagerPro/infra"
	"go-gin-productManagerPro/repositories"
	"go-gin-productManagerPro/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	// products := []models.Product{
	// 	{ID: 1, Name: "Product1", Price: 1000, Description: "Description1", SoldOut: false},
	// 	{ID: 2, Name: "Product2", Price: 2000, Description: "Description2", SoldOut: true},
	// 	{ID: 3, Name: "Product3", Price: 3000, Description: "Description3", SoldOut: false},
	// }

	// productRepository := repositories.NewProductMemoryRepository(products)
	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	r := gin.Default()
	// ルーターのグルーピング
	productRouter := r.Group("/products")
	productRouter.GET("", productController.FindAll)
	productRouter.GET("/:id", productController.FindById)
	productRouter.POST("", productController.Create)
	productRouter.PUT("/:id", productController.Update)
	productRouter.DELETE("/:id", productController.Delete)
	r.Run("localhost:8080")
}
