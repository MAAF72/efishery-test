package models

import (
	"golang.org/x/crypto/bcrypt"
)

// User user struct
type User struct {
	Base
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Hash        string `json:"-"`
}

// NewUser new user from request
func NewUser(request UserCreateRequest) User {
	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	return User{
		PhoneNumber: request.PhoneNumber,
		Name:        request.Name,
		Role:        request.Role,
		Hash:        string(hash),
	}
}
