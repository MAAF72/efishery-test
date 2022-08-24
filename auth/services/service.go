package services

import (
	"github.com/MAAF72/efishery-test/adapters"
	"github.com/MAAF72/efishery-test/repositories"
)

type service interface {
	UserService
}

// Instance service singleton
var Instance service

// Service service struct
type Service struct {
	repositories repositories.Repositories
}

// Init init services
func Init(adapters adapters.Adapters) {
	repositories := repositories.Init(adapters)
	Instance = Service{
		repositories: repositories,
	}
}
