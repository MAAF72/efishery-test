package repositories

import "github.com/MAAF72/efishery-test/adapters"

// Repositories repositories interface
type Repositories interface {
}

type repositories struct {
	adapters adapters.Adapters
}

// Init init repositories
func Init(adapters adapters.Adapters) Repositories {
	return repositories{
		adapters,
	}
}
