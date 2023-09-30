package product

import (
	"log"

	"github.com/nasrunrozikinmm/bank_sampah.git/src/config"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() []Product
	FindOne(id int) Product
	Save(product Product) (*Product, error)
	// Update(product Product) (*Product, error)
	// Delete(product Product) (*Product, error)
	// SoftDelete(product Product) (*Product, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db}
}

func (repo *ProductRepositoryImpl) FindAll() []Product {
	var products []Product
	_ = repo.db.Find(&products)
	return products
}
func (repo *ProductRepositoryImpl) FindOne(id int) Product {
	var product Product
	db, _ := config.DBSqlConnect()
	row := db.QueryRow("SELECT * FROM products WHERE id = ?", id)

	if err := row.Scan(
		&product.ID,
		&product.Title,
		&product.Category,
		&product.Author,
		&product.Deleted,
		&product.CreatedAt,
		&product.UpdatedAt,
	); err != nil {
		log.Println(err)

	}
	return product
}
func (repo *ProductRepositoryImpl) Save(product Product) (*Product, error) {
	result := repo.db.Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}
