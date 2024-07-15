package main

import (
	"go-gin-productManagerPro/controllers"
	"go-gin-productManagerPro/infra"
	"go-gin-productManagerPro/middlewares"
	"go-gin-productManagerPro/repositories"
	"go-gin-productManagerPro/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	// productRepository := repositories.NewProductMemoryRepository(products)
	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	authRrepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthSevice(authRrepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	r.Use(cors.Default())
	// ルーターのグルーピング
	productRouter := r.Group("/products")
	productRouterWithAuth := r.Group("/products", middlewares.AuthMiddlware(authService))
	authRouter := r.Group("/auth")

	productRouter.GET("", productController.FindAll)
	productRouterWithAuth.GET("/:id", productController.FindById)
	productRouterWithAuth.POST("", productController.Create)
	productRouterWithAuth.PUT("/:id", productController.Update)
	productRouterWithAuth.DELETE("/:id", productController.Delete)

	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	return r
}

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	r := setupRouter(db)
	r.Run("localhost:8080")
}

// products := []models.Product{
// 	{ID: 1, Name: "Product1", Price: 1000, Description: "Description1", SoldOut: false},
// 	{ID: 2, Name: "Product2", Price: 2000, Description: "Description2", SoldOut: true},
// 	{ID: 3, Name: "Product3", Price: 3000, Description: "Description3", SoldOut: false},
// }
