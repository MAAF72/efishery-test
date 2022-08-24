package repositories

import "github.com/MAAF72/efishery-test/adapters"

// Repositories repositories interface
type Repositories interface {
	UserRepository
}

type repositories struct {
	adapters adapters.Adapters
	DatabaseRepository
}

// Init init repositories
func Init(adapters adapters.Adapters) Repositories {
	return repositories{
		adapters,
		DatabaseRepository{db: adapters.Database},
	}
}
