package use_cases

import (
	"context"
	"github.com/google/uuid"
	"go-auth-api/internal/application/validators"
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/exceptions"
	"go-auth-api/internal/domain/types"
)

type findUserUseCaseImpl struct {
	userService adapters.UserService
}

func FindUserUseCaseConstructor(userService adapters.UserService) adapters.FindUserUseCaseAdapter {
	return &findUserUseCaseImpl{
		userService: userService,
	}
}

func (f findUserUseCaseImpl) ListUser(ctx context.Context, params dtos.QueryParams) (*[]types.User, exceptions.ErrorType) {
	var err exceptions.ErrorType

	err = validators.Validate(params, ctx)

	if err != nil {

		return nil, err
	}

	return f.userService.ListUser(ctx, params)

}

func (f findUserUseCaseImpl) FindUser(ctx context.Context, id uuid.UUID) (*types.User, exceptions.ErrorType) {
	user, err := f.userService.FindUserById(ctx, &id)

	if user == nil {
		return nil, exceptions.NotFoundException(ctx, "User not found for id: "+id.String())
	}

	return user, err
}
