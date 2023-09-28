package repositories

import (
	"github.com/iyunrozikin26/bank_sampah.git/src/domains/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() []models.Product
	// FindOne(id int) models.Product
	Save(product models.Product) (*models.Product, error)
	// Update(product models.Product) (*models.Product, error)
	// Delete(product models.Product) (*models.Product, error)
	// SoftDelete(product models.Product) (*models.Product, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db}
}

func (revo *ProductRepositoryImpl) FindAll() []models.Product {
	var products []models.Product
	_ = revo.db.Find(&products)
	return products
}
func (revo *ProductRepositoryImpl) Save(product models.Product) (*models.Product, error) {
	result := revo.db.Create(&product)
	if result != nil {
		return nil, result.Error
	}
	return &product, nil
}
