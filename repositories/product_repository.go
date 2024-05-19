package repositories

import "go-gin-productManagerPro/models"

type IProductRepository interface {
	FindAll() (*[]models.Product, error)
}

type ProductMemoryRepository struct {
	products []models.Product
}

func NewProductMemoryRepository(products []models.Product) IProductRepository {
	return &ProductMemoryRepository{products: products}
}

func (r *ProductMemoryRepository) FindAll() (*[]models.Product, error) {
	return &r.products, nil
}
