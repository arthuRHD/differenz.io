package api

import (
	"github.com/go-playground/validator/v10"
)

type FieldValidator struct {
	validator.Validate
}

type Violations struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewFieldValidator() *FieldValidator {
	return &FieldValidator{*validator.New()}
}

func (v *FieldValidator) Check(object interface{}) error {
	return v.Struct(object)
}

func ToViolations(err error) []Violations {
	var violations []Violations
	for _, e := range err.(validator.ValidationErrors) {
		violations = append(violations, Violations{Field: e.Field(), Message: e.Tag()})
	}
	return violations
}
