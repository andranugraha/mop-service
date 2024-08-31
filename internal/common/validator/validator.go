package validator

import (
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

var (
	validate     *validator.Validate
	validateOnce sync.Once
)

func Validate(data interface{}) error {
	validateOnce.Do(func() {
		validate = validator.New()
	})

	validationErrors := make([]string, 0)
	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.Field(),
				err.Value(),
				err.Tag(),
			))
		}
		return errors.New(fmt.Sprintf("Validation Error: %v", validationErrors))
	}

	return nil
}
