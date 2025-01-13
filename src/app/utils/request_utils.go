package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	apperrors "sm.com/m/src/app/app_errors"
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

func FormatAndSendRequiredFieldsError(err error, c *gin.Context) {
	var verr validator.ValidationErrors
	if errors.As(err, &verr) {
		c.JSON(http.StatusBadRequest, ResponseError(
			apperrors.ErrMissingFields,
			DescriptiveError(verr),
		))
	} else {
		c.JSON(http.StatusBadRequest, ResponseError(
			apperrors.ErrInvalidRequest,
			err.Error(),
		))
	}
}
