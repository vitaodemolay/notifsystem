package internalerrors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator"
)

func ValidateStruct(obj any) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)

	msgs := mapErrorsToArray(validationErrors)

	return errors.New(strings.Join(msgs, ";"))
}

func mapErrorsToArray(valErrors validator.ValidationErrors) []string {
	var errors []string
	fieldErrorResolver := newStrategyFieldErrorResolver()

	for _, fieldError := range valErrors {
		errors = append(errors, fieldErrorResolver.Resolve(fieldError))
	}
	return errors
}

type strategyFieldErrorResolver struct {
	resolver map[string]func(validator.FieldError) string
}

func newStrategyFieldErrorResolver() *strategyFieldErrorResolver {
	return &strategyFieldErrorResolver{
		resolver: map[string]func(validator.FieldError) string{
			"required": requiredFieldErrorResolver,
			"max":      maxFieldErrorResolver,
			"min":      minFieldErrorResolver,
			"email":    emailFieldErrorResolver,
		},
	}
}

func (s *strategyFieldErrorResolver) Resolve(ferr validator.FieldError) string {
	if resolver, ok := s.resolver[ferr.Tag()]; ok {
		return resolver(ferr)
	}
	return ""
}

func requiredFieldErrorResolver(ferr validator.FieldError) string {
	field := strings.ToLower(ferr.StructField())
	return field + " is required"
}

func maxFieldErrorResolver(ferr validator.FieldError) string {
	field := strings.ToLower(ferr.StructField())
	return field + " is required with max " + ferr.Param()
}

func minFieldErrorResolver(ferr validator.FieldError) string {
	field := strings.ToLower(ferr.StructField())
	return field + " is required with min " + ferr.Param()
}

func emailFieldErrorResolver(ferr validator.FieldError) string {
	field := strings.ToLower(ferr.StructField())
	return field + " is invalid"
}
