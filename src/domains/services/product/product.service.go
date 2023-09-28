package services

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iyunrozikin26/bank_sampah.git/src/domains/models"
	"github.com/iyunrozikin26/bank_sampah.git/src/domains/repositories"
)

// method contracts
type ProductService interface {
	GetAll() []models.Product
	GetByID(id int) models.Product
	Create(ctx *gin.Context) (*models.Product, error)
	// Update(ctx *gin.Context) (*models.Product, error)
	// Delete(ctx *gin.Context) (*models.Product, error)
}

type ProductServiceImpl struct {
	productRepository repositories.ProductRepository
}

// constructor
func NewProductService(productRepository repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepository}
}

// ps = product service
func (serv *ProductServiceImpl) GetAll() []models.Product {
	return serv.productRepository.FindAll()
}
func (serv *ProductServiceImpl) GetByID(id int) models.Product {
	return serv.productRepository.FindOne(id)
}

func (serv *ProductServiceImpl) Create(ctx *gin.Context) (*models.Product, error) {
	var input CreateProductDTO
	// binding to JSON
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}
	// validasi input
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return nil, err
	}
	// mengisi struck dengan input dari client
	productPayload := models.Product{
		Title:     input.Title,
		Category:  input.Category,
		Author:    input.Author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := serv.productRepository.Save(productPayload)
	log.Println(result, "qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq")
	log.Println(err, "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	if err != nil {
		return nil, err
	}
	return result, nil
}
