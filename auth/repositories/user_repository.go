package repositories

import (
	"github.com/MAAF72/efishery-test/models"
)

// UserRepository user repository interface
type UserRepository interface {
	SaveUser(request models.User) (id string, err error)
	FindUserByID(id string) (user models.User, err error)
	FindUserByPhoneNumber(phoneNumber string) (user models.User, err error)
}

// SaveUser save user
func (repo DatabaseRepository) SaveUser(user models.User) (id string, err error) {
	err = repo.db.Create(&user).Error

	id = user.ID

	return
}

// FindUserByID find user by id
func (repo DatabaseRepository) FindUserByID(id string) (user models.User, err error) {
	err = repo.db.Take(&user, "id = ?", id).Error

	return
}

// FindUserByPhoneNumber find user by phone number
func (repo DatabaseRepository) FindUserByPhoneNumber(phoneNumber string) (user models.User, err error) {
	err = repo.db.Take(&user, "phone_number = ?", phoneNumber).Error

	return
}
