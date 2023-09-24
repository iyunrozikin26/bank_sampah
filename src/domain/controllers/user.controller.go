package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userServ "github.com/iyunrozikin26/bank_sampah.git/src/domain/services/user"
	"github.com/iyunrozikin26/bank_sampah.git/src/helpers"
)

type UserController struct {
	userService userServ.UserService
	ctx         *gin.Context
}

func NewUserController(userService userServ.UserService, ctx *gin.Context) UserController {
	return UserController{userService, ctx}
}

func (uc *UserController) Index(ctx *gin.Context) {
	data := uc.userService.GetAll()
	var message string
	message = "success get all data"
	helpers.ResponseJSON(ctx, http.StatusOK, message, data)
}
func (uc *UserController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := uc.userService.GetByID(id)
	var message string

	message = "success get data by ID"
	helpers.ResponseJSON(ctx, http.StatusOK, message, data)
}
func (uc *UserController) Create(ctx *gin.Context) {
	data, err := uc.userService.Create(ctx)
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
func (uc *UserController) Update(ctx *gin.Context) {
	data, err := uc.userService.Update(ctx)
	var message string

	if err != nil {
		message = "Internal server error"
		helpers.ResponseJSON(ctx, http.StatusInternalServerError, message, err)
		ctx.Abort()
		return
	}

	message = "Update successfully"
	helpers.ResponseJSON(ctx, http.StatusCreated, message, data)
}
func (uc *UserController) Delete(ctx *gin.Context) {
	data, err := uc.userService.Delete(ctx)
	var message string

	if err != nil {
		message = "Internal server error"
		helpers.ResponseJSON(ctx, http.StatusInternalServerError, message, err)
		ctx.Abort()
		return
	}
	message = "Delete successfully"
	helpers.ResponseJSON(ctx, http.StatusOK, message, data)
}
