package product

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dto "github.com/nasrunrozikinmm/bank_sampah.git/src/modules/product/dto"
)

// method contracts
type ProductService interface {
	GetAll() []Product
	GetByID(id int) Product
	Create(ctx *gin.Context) (*Product, error)
	// Update(ctx *gin.Context) (*Product, error)
	// Delete(ctx *gin.Context) (*Product, error)
}

type ProductServiceImpl struct {
	productRepository ProductRepository
}

// constructor
func NewProductService(productRepository ProductRepository) ProductService {
	return &ProductServiceImpl{productRepository}
}

// ps = product service
func (serv *ProductServiceImpl) GetAll() []Product {
	return serv.productRepository.FindAll()
}
func (serv *ProductServiceImpl) GetByID(id int) Product {
	return serv.productRepository.FindOne(id)
}

func (serv *ProductServiceImpl) Create(ctx *gin.Context) (*Product, error) {
	var input dto.CreateProductDTO
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
	productPayload := Product{
		Title:     input.Title,
		Category:  input.Category,
		Author:    input.Author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := serv.productRepository.Save(productPayload)
	if err != nil {
		return nil, err
	}
	return result, nil
}
