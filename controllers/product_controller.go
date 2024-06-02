package controllers

import (
	"go-gin-productManagerPro/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IProductController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
}

type ProductController struct {
	service services.IProductService
}

func NewProductController(service services.IProductService) IProductController {
	return &ProductController{service: service}
}

func (c *ProductController) FindAll(ctx *gin.Context) {
	products, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return 
	}
	ctx.JSON(http.StatusOK, gin.H{"date": products})
}

func (c *ProductController) FindById(ctx *gin.Context) {
	productId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return 
	}
	product, err := c.service.FindById(uint(productId))
	if err != nil {
		if err.Error() == "products not found"{
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON((http.StatusOK), gin.H{"date": product})
}