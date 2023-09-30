package user

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dto "github.com/nasrunrozikinmm/bank_sampah.git/src/modules/user/dto"
)

type UserService interface {
	GetAll() []User
	GetByID(id int) User
	Create(ctx *gin.Context) (*User, error) // ctx *gin.Context untuk menangkap params, query, dll dari input user
	Update(ctx *gin.Context) (*User, error)
	Delete(ctx *gin.Context) (*User, error)
}

type UserServiceImpl struct {
	userRepository UserRepository
}

// constructor
func NewUserService(userRepository UserRepository) UserService {
	return &UserServiceImpl{userRepository}
}

// us = user service
func (us *UserServiceImpl) GetAll() []User {
	return us.userRepository.FindAll()
}

func (us *UserServiceImpl) GetByID(id int) User {
	return us.userRepository.FindOne(id)
}

func (us *UserServiceImpl) Create(ctx *gin.Context) (*User, error) {
	var userInput dto.CreateUserDTO
	// binding to JSON
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		return nil, err
	}
	// validasi input
	validate := validator.New()
	err := validate.Struct(userInput)
	if err != nil {
		return nil, err
	}
	// mengisi struck dengan input dari client
	userPayload := User{
		Name:      userInput.Name,
		Email:     userInput.Email,
		Password:  userInput.Password,
		Role:      userInput.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := us.userRepository.Save(userPayload)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (us *UserServiceImpl) Update(ctx *gin.Context) (*User, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var userInput dto.UpdateUserDTO

	// binding to JSON
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		return nil, err
	}
	// validasi input
	validate := validator.New()
	err := validate.Struct(userInput)
	if err != nil {
		return nil, err
	}
	// mengisi struck dengan input dari client
	userPayload := User{
		ID:        int64(id),
		Name:      userInput.Name,
		Email:     userInput.Email,
		Password:  userInput.Password,
		UpdatedAt: time.Now(),
	}
	result, err := us.userRepository.Update(userPayload)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (us *UserServiceImpl) Delete(ctx *gin.Context) (*User, error) {
	id, _ := strconv.Atoi(ctx.Param("id")) // untuk mendapatkan id dan Atoi to mengconvert to ParseInt
	user := User{
		ID: int64(id),
	}
	result, err := us.userRepository.Delete(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
