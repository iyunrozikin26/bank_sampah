package controllers

import (
	"net/http"

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
