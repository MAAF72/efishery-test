package database

import "github.com/MAAF72/efishery-test/models"

// Migrate migrate the models
func Migrate(db *Database) (err error) {
	modelList := []interface{}{
		models.User{},
	}

	for _, model := range modelList {
		err = db.AutoMigrate(model)
		if err != nil {
			return
		}
	}

	return
}
