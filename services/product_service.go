package services

import (
	"go-gin-productManagerPro/dto"
	"go-gin-productManagerPro/models"
	"go-gin-productManagerPro/repositories"
)

type IProductService interface {
	FindAll() (*[]models.Product, error)
	FindById(productId uint) (*models.Product, error)
	Create(createProductInput dto.CreateItemInput) (*models.Product, error)
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

func (s *ProductService) Create(createProductInput dto.CreateItemInput) (*models.Product, error) {
	newProduct := models.Product{
		Name:        createProductInput.Name,
		Price:       createProductInput.Price,
		Description: createProductInput.Description,
		SoldOut: false,
	}
	return s.repository.Create(newProduct)
}
