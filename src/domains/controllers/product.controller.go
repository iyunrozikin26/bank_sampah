package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	services "github.com/iyunrozikin26/bank_sampah.git/src/domains/services/product"
	"github.com/iyunrozikin26/bank_sampah.git/src/helpers"
)

type ProductController struct {
	productService services.ProductService
	ctx            *gin.Context
}

func NewProductController(productService services.ProductService, ctx *gin.Context) ProductController {
	return ProductController{productService, ctx}
}

func (pc *ProductController) IndexProduct(ctx *gin.Context) {
	data := pc.productService.GetAll()
	var message string
	message = "success get all data"
	helpers.ResponseJSON(ctx, http.StatusOK, message, data)
}
func (pc *ProductController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := pc.productService.GetByID(id)
	var message string

	message = "success get data by ID"
	helpers.ResponseJSON(ctx, http.StatusOK, message, data)
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	data, err := pc.productService.Create(ctx)
	var message string

	if err != nil {
		message = "Internal server error"
		helpers.ResponseJSON(ctx, http.StatusInternalServerError, message, err)
		ctx.Abort()
		return
	}
	message = "Create successfully"
	helpers.ResponseJSON(ctx, http.StatusCreated, message, data)

}
