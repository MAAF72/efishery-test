package services

import (
	"errors"

	"github.com/MAAF72/efishery-test/models"
	gonanoid "github.com/matoous/go-nanoid"
)

// UserService user service interface
type UserService interface {
	CreateUser(request models.UserCreateRequest) (res models.UserCreateResponse, err error)
	GetUserByID(id string) (res models.User, err error)
	GetUserByPhoneNumber(phoneNumber string) (res models.User, err error)
}

// CreateUser create
func (service Service) CreateUser(request models.UserCreateRequest) (res models.UserCreateResponse, err error) {
	// find another user with same phone number
	userWithSamePhoneNumber, _ := service.repositories.FindUserByPhoneNumber(request.PhoneNumber)

	if userWithSamePhoneNumber.Base.ID != "" {
		err = errors.New("phone number already registered")
		return
	}

	// generate random password (4 characters)
	generatedPassword, _ := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyz", 4)

	request.Password = generatedPassword
	newUser := models.NewUser(request)

	id, err := service.repositories.SaveUser(newUser)
	if err != nil {
		return
	}

	user, err := service.repositories.FindUserByID(id)

	res = models.UserCreateResponse{
		Base:        user.Base,
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
		Role:        user.Role,
		Password:    generatedPassword,
	}

	return
}

// GetUserByID get user by id
func (service Service) GetUserByID(id string) (res models.User, err error) {
	return service.repositories.FindUserByID(id)
}

// GetUserByPhoneNumber get user by phone number
func (service Service) GetUserByPhoneNumber(phoneNumber string) (res models.User, err error) {
	return service.repositories.FindUserByPhoneNumber(phoneNumber)
}
