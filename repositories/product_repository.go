package repositories

import (
	"errors"
	"go-gin-productManagerPro/models"
)

type IProductRepository interface {
	FindAll() (*[]models.Product, error)
	FindById(itemId uint) (*models.Product, error)
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

func (r *ProductMemoryRepository) FindById(itemId uint) (*models.Product, error) {
	for _, v := range r.products {
		if v.ID == itemId {
			return &v, nil
		}
	}
	return nil, errors.New("products not found")
}
