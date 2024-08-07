package controllers

import (
	"go-gin-productManagerPro/dto"
	"go-gin-productManagerPro/models"
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
	Delete(ctx *gin.Context)
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
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userId := user.(*models.User).ID

	productId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	product, err := c.service.FindById(uint(productId), userId)
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
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// *models.User 型に変換
	userId := user.(*models.User).ID

	var input dto.CreateProductInput

	// respond with a 400 BadRequest status and provide a descriptive error message.
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct, err := c.service.Create(input, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newProduct})
}

func (c *ProductController) Update(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userId := user.(*models.User).ID

	productId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	var input dto.UpdateProductInput
	/**
	 * Ginフレームワークで提供されているメソッドで、
	 * HTTPリクエストのボディからJSONデータを読み取り
	 * 指定されたGoの構造体にバインドする機能
	 **/
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateProduct, err := c.service.Update(uint(productId), userId, input)
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

func (c *ProductController) Delete(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userId := user.(*models.User).ID

	productId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	err = c.service.Delete(uint(productId), userId)
	if err != nil {
		if err.Error() == "product not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.Status(http.StatusOK)
}
