package controllers

import (
	"go-gin-productManagerPro/dto"
	"go-gin-productManagerPro/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IProductController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
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
		if err.Error() == "products not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON((http.StatusOK), gin.H{"date": product})
}

func (c *ProductController) Create(ctx *gin.Context) {
	var input dto.CreateProductInput

	// respond with a 400 BadRequest status and provide a descriptive error message.
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newProduct})
}

func (c *ProductController) Update(ctx *gin.Context) {
	productId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	var input dto.UpdateProductInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateProduct, err := c.service.Update(uint(productId), input)
	if err != nil {
		if err.Error() == "products not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"date": updateProduct})
}

/* 知識の補足*/
/*
ctx.JSONメソッド： Ginフレームワークの一部
返り値として認識される理由は、
この行がHTTPレスポンスを生成しクライアントに送信するためです。
指定されたHTTPステータスコードと共にJSON形式のデータをクライアントに返す
*/
