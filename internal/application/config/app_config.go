package config

import (
	"go-auth-api/internal/domain/adapters"
	"gorm.io/gorm"
	"sync"
)

var injectorInit sync.Once
var injector *AppConfiguration

type AppConfiguration struct {
	UserController adapters.UsersController
	db             *gorm.DB
}

func GetInjector() *AppConfiguration {
	if injector == nil {
		injectorInit.Do(func() {
			injector = &AppConfiguration{}
		})
	}
	return injector
}

func (ac *AppConfiguration) InjectDependencies() *AppConfiguration {

	return ac
}
