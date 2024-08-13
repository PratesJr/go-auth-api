package validators

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"go-auth-api/internal/domain/exceptions"
)

func Validate(object interface{}, ctx context.Context) exceptions.ErrorType {
	validate := validator.New()
	err := validate.Struct(object)
	if err == nil {
		return nil
	}
	validationErr := err.(validator.ValidationErrors)[0]
	var ex []error

	switch validationErr.Tag() {

	case "required":
		ex = append(ex, errors.New(validationErr.StructField()+" is required"))

	case "min":
		ex = append(ex, errors.New(validationErr.StructField()+" is required with min "+validationErr.Param()))

	case "max":
		ex = append(ex, errors.New(validationErr.StructField()+" is required with max "+validationErr.Param()))

	case "email":
		ex = append(ex, errors.New(validationErr.StructField()+" is not valid"))
	}

	if len(ex) > 0 {
		return exceptions.UnprocessableEntity(ctx, "invalid payload", ex)
	}

	return nil
}
