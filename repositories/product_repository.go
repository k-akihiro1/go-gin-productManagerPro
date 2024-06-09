package repositories

import (
	"errors"
	"go-gin-productManagerPro/models"
)

type IProductRepository interface {
	FindAll() (*[]models.Product, error)
	FindById(productId uint) (*models.Product, error)
	Create(newProduct models.Product) (*models.Product, error)
	Update(updateProduct models.Product)(*models.Product, error)
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

func (r *ProductMemoryRepository) FindById(productId uint) (*models.Product, error) {
	for _, v := range r.products {
		if v.ID == productId {
			return &v, nil
		}
	}
	return nil, errors.New("products not found")
}

func (r *ProductMemoryRepository) Create(newProduct models.Product) (*models.Product, error) {
	newProduct.ID = uint(len(r.products) + 1)
	r.products = append(r.products, newProduct)
	return &newProduct, nil
}

/* 知識の補足*/
/*
return &newProduct, nilで返される&newProductは、
newProduct変数のアドレス、つまりメモリ上の位置を指すポインタです。
newProductはmodels.Product型の変数であり、
&演算子を使用することでその変数のアドレスを取得し、
*models.Product型（models.Productのポインタ型）の値として返している
*/

func (r *ProductMemoryRepository) Update(updateProduct models.Product)(*models.Product, error){
	for i, v := range r.products {
		if v.ID == updateProduct.ID{
			r.products[i] = updateProduct
			return &r.products[i], nil	
		}
	}
	return nil, errors.New("unexpected error")
}