package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// APIValidationError - wraps API validation errors
type APIValidationError struct {
	ActualTag string `json:"tag"`
	Namespace string `json:"field"`
	Kind      string `json:"kind"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	Param     string `json:"param"`
}

// WrapAPIValidationErrors - function to wrap errors in struct
func WrapAPIValidationErrors(errs validator.ValidationErrors) []APIValidationError {
	validationErrors := make([]APIValidationError, 0, len(errs))
	for _, validationErr := range errs {
		validationErrors = append(validationErrors, APIValidationError{
			ActualTag: validationErr.ActualTag(),
			Namespace: validationErr.Namespace(),
			Kind:      validationErr.Kind().String(),
			Type:      validationErr.Type().String(),
			Value:     fmt.Sprintf("%v", validationErr.Value()),
			Param:     validationErr.Param(),
		})
	}

	return validationErrors
}
