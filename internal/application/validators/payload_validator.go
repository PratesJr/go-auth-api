package validators

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-auth-api/internal/domain/exceptions"
)

func Validate(object interface{}, ctx context.Context) exceptions.ErrorType {
	validate := validator.New()
	err := validate.Struct(object)
	fmt.Println(err)
	if err == nil {
		return nil
	}

	var validationErr validator.ValidationErrors
	errors.As(err, &validationErr)

	var ex []exceptions.ErrorDetails

	for _, e := range validationErr {
		switch e.Tag() {

		case "required":
			ex = append(ex, exceptions.NewErrorDetail(e.StructField(), " is required"))

		case "min":
			ex = append(ex, exceptions.NewErrorDetail(e.StructField(), " is required with min "+e.Param()))

		case "max":
			ex = append(ex, exceptions.NewErrorDetail(e.StructField(), " is required with max "+e.Param()))

		case "email":
			ex = append(ex, exceptions.NewErrorDetail(e.StructField(), " is not valid"))
		}
	}

	if len(ex) > 0 {
		return exceptions.UnprocessableEntity(ctx, "invalid payload", ex)
	}

	return nil
}
