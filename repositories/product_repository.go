package repositories

import (
	"errors"
	"go-gin-productManagerPro/models"

	"gorm.io/gorm"
)

type IProductRepository interface {
	FindAll() (*[]models.Product, error)
	FindById(productId uint, userId uint) (*models.Product, error)
	Create(newProduct models.Product) (*models.Product, error)
	Update(updateProduct models.Product) (*models.Product, error)
	Delete(productId uint, userId uint) error
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

func (r *ProductMemoryRepository) FindById(productId uint, userId uint) (*models.Product, error) {
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

func (r *ProductMemoryRepository) Update(updateProduct models.Product) (*models.Product, error) {
	for i, v := range r.products {
		if v.ID == updateProduct.ID {
			r.products[i] = updateProduct
			return &r.products[i], nil
		}
	}
	return nil, errors.New("unexpected error")
}

func (r *ProductMemoryRepository) Delete(productId uint, userId uint) error {
	for i, v := range r.products {
		if v.ID == productId {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}

/* 知識の補足*/
/*
スライス式の基本形
スライス[開始インデックス:終了インデックス]

スライス[:i] は、スライスの最初の要素から i-1 番目の要素まで
スライス[i+1:] は、スライスの i+1 畖目の要素から最後の要素まで
*/

// DB接続用
type productRepository struct {
	db *gorm.DB
}

// Create implements IProductRepository.
func (r *productRepository) Create(newProduct models.Product) (*models.Product, error) {
	result := r.db.Create(&newProduct)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newProduct, nil
}

// Delete implements IProductRepository.
func (p *productRepository) Delete(productId uint, userId uint) error {
	deleteItem, err := p.FindById(productId, userId);
	if err != nil {
		return err
	}

	// result := p.db.Unscoped().Delete(&deleteItem) // 物理削除
	result := p.db.Delete(&deleteItem) // 論理削除
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindAll implements IProductRepository.
func (p *productRepository) FindAll() (*[]models.Product, error) {
	var products []models.Product
	result := p.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return &products, nil
}

// FindById implements IProductRepository.
func (p *productRepository) FindById(productId uint, userId uint) (*models.Product, error) {
	var product models.Product
	result := p.db.First(&product, "id = ? AND user_id = ?", productId, userId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("product not found")
		}
		return nil, result.Error
	}
	return &product, nil
}

// Update implements IProductRepository.
func (p *productRepository) Update(updateProduct models.Product) (*models.Product, error) {
	result := p.db.Save(&updateProduct)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateProduct, nil
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{db: db}
}
