package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	models "github.com/iyunrozikin26/bank_sampah.git/src/domain/models"
	userRepo "github.com/iyunrozikin26/bank_sampah.git/src/domain/repositories"
)

type UserService interface {
	GetAll() []models.User
	GetByID(id int) models.User
	Create(ctx *gin.Context) (*models.User, error) // ctx *gin.Context untuk menangkap params, query, dll dari input user
	Update(ctx *gin.Context) (*models.User, error)
	Delete(ctx *gin.Context) (*models.User, error)
}

type UserServiceImpl struct {
	userRepository userRepo.UserRepository
}

// constructor

func NewUserService(userRepository userRepo.UserRepository) UserService {
	return &UserServiceImpl{userRepository}
}

func (us *UserServiceImpl) GetAll() []models.User {
	return us.userRepository.FindAll()
}

func (us *UserServiceImpl) GetByID(id int) models.User {
	return us.userRepository.FindOne(id)
}

func (us *UserServiceImpl) Create(ctx *gin.Context) (*models.User, error) {
	var userInput CreateUserDTO
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
	userPayload := models.User{
		Name:      userInput.Name,
		Email:     userInput.Email,
		Password:  userInput.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	fmt.Println(userPayload, "qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq")
	result, err := us.userRepository.Save(userPayload)
	fmt.Println(result, "qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq")
	fmt.Println(err, "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (us *UserServiceImpl) Update(ctx *gin.Context) (*models.User, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var userInput UpdateUserDTO

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
	fmt.Println(userInput, "userInputuserInputuserInputuserInputuserInput")
	// mengisi struck dengan input dari client
	userPayload := models.User{
		ID:        int64(id),
		Name:      userInput.Name,
		Email:     userInput.Email,
		Password:  userInput.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := us.userRepository.Update(userPayload)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (us *UserServiceImpl) Delete(ctx *gin.Context) (*models.User, error) {
	id, _ := strconv.Atoi(ctx.Param("id")) // untuk mendapatkan id dan Atoi to mengconvert to ParseInt
	user := models.User{
		ID: int64(id),
	}
	result, err := us.userRepository.Delete(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
