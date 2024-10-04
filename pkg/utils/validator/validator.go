package util_validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// validate with paging
var validate *validator.Validate

func init() {
	validate = validator.New()
}
func ValidatePagin(s interface{}) error {
	return validate.Struct(s)
}

// normal validate
type ValidateResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(value interface{}) []*ValidateResponse {
	var errors []*ValidateResponse
	validate := validator.New()
	err := validate.Struct(value)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidateResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateStructPaging(value interface{}) ([]string, error) {
	// fmt.Println("hello")
	err := validate.Struct(value)

	// fmt.Println("Error = ", err)
	if err != nil {
		// Check if the error is of type ValidationErrors
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			// collect formatted errors
			var errs []string
			for _, validationErr := range validationErrors {
				errs = append(
					errs, fmt.Sprintf(
						"Field '%s' failed on the '%s' tag",
						validationErr.Field(),
						validationErr.Tag(),
					),
				)
			}
			return errs, err
		}
	}
	return nil, nil
}

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
