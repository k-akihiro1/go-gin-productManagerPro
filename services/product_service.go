package services

import (
	"go-gin-productManagerPro/dto"
	"go-gin-productManagerPro/models"
	"go-gin-productManagerPro/repositories"
)

type IProductService interface {
	FindAll() (*[]models.Product, error)
	FindById(productId uint) (*models.Product, error)
	Create(createProductInput dto.CreateProductInput) (*models.Product, error)
	Update(productId uint, updateproductInput dto.UpdateProductInput) (*models.Product, error)
	Delete(productId uint) error
}

type ProductService struct {
	repository repositories.IProductRepository
}

func NewProductService(repository repositories.IProductRepository) IProductService {
	return &ProductService{repository: repository}
}

func (s *ProductService) FindAll() (*[]models.Product, error) {
	return s.repository.FindAll()
}

func (s *ProductService) FindById(productId uint) (*models.Product, error) {
	return s.repository.FindById(productId)
}

func (s *ProductService) Create(createProductInput dto.CreateProductInput) (*models.Product, error) {
	newProduct := models.Product{
		Name:        createProductInput.Name,
		Price:       createProductInput.Price,
		Description: createProductInput.Description,
		SoldOut:     false,
	}
	return s.repository.Create(newProduct)
}

func (s *ProductService) Update(productId uint, updateProductInput dto.UpdateProductInput) (*models.Product, error) {
	targetProduct, err := s.FindById(productId)
	if err != nil {
		return nil, err
	}

	if updateProductInput.Name != nil {
		targetProduct.Name = *updateProductInput.Name
	}
	if updateProductInput.Price != nil {
		targetProduct.Price = *updateProductInput.Price
	}
	if updateProductInput.Description != nil {
		targetProduct.Description = *updateProductInput.Description
	}
	if updateProductInput.SoldOut != nil {
		targetProduct.SoldOut = *updateProductInput.SoldOut
	}
	return s.repository.Update(*targetProduct)
}

func (s *ProductService) Delete(ProductId uint) error {
	return s.repository.Delete(ProductId)
}
