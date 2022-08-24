package repositories

import (
	"github.com/MAAF72/efishery-test/adapters/database"
)

// DatabaseRepository database repository
type DatabaseRepository struct {
	db *database.Database
}
