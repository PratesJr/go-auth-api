package config

import (
	"gorm.io/gorm"
	"sync"
)

var injectorInit sync.Once
var injector *AppConfiguration

type AppConfiguration struct {
	db *gorm.DB
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
