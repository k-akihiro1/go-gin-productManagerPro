package controllers

import (
	"go-gin-productManagerPro/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IProductController interface {
	FindAll(ctx *gin.Context)
}

type ProductController struct {
	service services.IProductService
}

func NewProductController(service services.IProductService) IProductController {
	return &ProductController{service: service}
}

func (c *ProductController) FindAll(ctx *gin.Context) {
	iterms, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return 
	}
	ctx.JSON(http.StatusOK, gin.H{"date": iterms})
}