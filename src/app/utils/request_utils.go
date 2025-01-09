package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func DescriptiveError(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)
	for _, field := range verr {
		tag := field.ActualTag()
		if field.Param() != "" {
			tag = fmt.Sprintf("%s=%s", tag, field.Param())
		}
		errs[field.Field()] = tag + ", " + field.Type().String()
	}
	return errs
}
