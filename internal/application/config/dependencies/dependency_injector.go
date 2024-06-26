package dependencies

import (
	"go-auth-api/internal/application/controllers"
	"go-auth-api/internal/application/database"
	"go-auth-api/internal/domain/adapters"
	use_cases "go-auth-api/internal/domain/use-cases"
	"go-auth-api/internal/integration/models"
	"go-auth-api/internal/integration/persistences"
	"go-auth-api/internal/integration/repositories"
	"gorm.io/gorm"
	"sync"
)

var injectorInit sync.Once
var injector *AppConfiguration

type AppConfiguration struct {
	db              *gorm.DB
	UserRepository  adapters.UserRepository
	UserPersistence adapters.UserPersistence
	UserUseCase     adapters.UserUseCase
	UserController  adapters.UsersController
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

	if ac.db == nil {
		ac.db = database.DbConnection()
	}

	if ac.UserRepository == nil {
		ac.UserRepository = repositories.UserRepositoryConstructor(*ac.db, models.UsersModel{})
	}

	if ac.UserPersistence == nil {
		ac.UserPersistence = persistences.UserPersistenceConstructor(ac.UserRepository)
	}

	if ac.UserUseCase == nil {
		ac.UserUseCase = use_cases.UserUseCaseConstructor(ac.UserPersistence)
	}

	if ac.UserController == nil {
		ac.UserController = controllers.UserControllerConstructor(ac.UserUseCase)
	}
	return ac
}
