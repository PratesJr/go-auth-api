package dependencies

import (
	"go-auth-api/internal/application/controllers"
	"go-auth-api/internal/application/database"
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/services"
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
	db                *gorm.DB
	userRepository    adapters.UserRepository
	userPersistence   adapters.UserPersistence
	userService       adapters.UserService
	createUserUseCase adapters.CreateUserUseCase
	updateUserUseCase adapters.UpdateUserUseCaseAdapter
	UserController    adapters.UsersController
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

	if ac.userRepository == nil {
		ac.userRepository = repositories.UserRepositoryConstructor(*ac.db, models.UsersModel{})
	}

	if ac.userPersistence == nil {
		ac.userPersistence = persistences.UserPersistenceConstructor(ac.userRepository)
	}

	if ac.userService == nil {
		ac.userService = services.UserServiceConstructor(ac.userPersistence)
	}

	if ac.createUserUseCase == nil {
		ac.createUserUseCase = use_cases.CreateUserUseCaseConstructor(ac.userService)
	}

	if ac.updateUserUseCase == nil {
		ac.updateUserUseCase = use_cases.UpdateUserUseCaseConstructor(ac.userService)
	}

	if ac.UserController == nil {

		ac.UserController = controllers.UserControllerConstructor(
			ac.createUserUseCase,
			ac.updateUserUseCase,
		)
	}

	return ac
}
