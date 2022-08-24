package adapters

import (
	"log"

	"github.com/MAAF72/efishery-test/adapters/database"
)

// Adapters adapters struct
type Adapters struct {
	Database *database.Database
}

// Init init all adapters
func Init() Adapters {
	db, err := database.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = database.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	return Adapters{
		Database: db,
	}
}
