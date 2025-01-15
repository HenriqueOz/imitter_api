package utilstest

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/utils"
)

type ValidateInput struct {
	Name string    `validate:"required"`
	Age  int       `validate:"required"`
	Date time.Time `validate:"required"`
}

func Test_DescriptiveError(t *testing.T) {
	var validate *validator.Validate

	setUp := func() {
		validate = validator.New()
	}

	t.Run("Missing all required fields", func(tt *testing.T) {
		setUp()

		input := ValidateInput{}
		err := validate.Struct(input)

		result := utils.DescriptiveError(err.(validator.ValidationErrors))

		assert.Equal(tt, 3, len(result), "Result must have 3 fields")
		assert.Equal(tt, "required, string", result["Name"], "Name field message must match the expect")
		assert.Equal(tt, "required, int", result["Age"], "Age field message must match the expect")
		assert.Equal(tt, "required, time.Time", result["Date"], "Date field message must match the expect")
	})

	t.Run("Missing one required fields", func(tt *testing.T) {
		setUp()

		input := ValidateInput{
			Name: "some name",
			Date: time.Now(),
		}
		err := validate.Struct(input)

		result := utils.DescriptiveError(err.(validator.ValidationErrors))

		assert.Equal(tt, 1, len(result), "Result must have 1 field")
		assert.Equal(tt, "required, int", result["Age"], "Age field message must match the expect")
	})

	t.Run("Not missing required fields", func(tt *testing.T) {
		setUp()

		input := ValidateInput{
			Name: "some name",
			Date: time.Now(),
		}
		err := validate.Struct(input)

		result := utils.DescriptiveError(err.(validator.ValidationErrors))

		assert.Equal(tt, 1, len(result), "Result must have 1 field")
		assert.Equal(tt, "required, int", result["Age"], "Age field message must match the expect")
	})
}

func Test_FormatAndSendRequiredFieldsError(t *testing.T) {
	var w *httptest.ResponseRecorder
	var ctx *gin.Context
	var validate *validator.Validate

	setUp := func() {
		gin.SetMode(gin.TestMode)

		validate = validator.New()
		w = httptest.NewRecorder()
		ctx, _ = gin.CreateTestContext(w)
	}

	t.Run("Should return a missing fields error when success on validate struct",
		func(tt *testing.T) {
			setUp()

			input := ValidateInput{
				Name: "some name",
			}
			err := validate.Struct(input)
			expectedBody := map[string]interface{}{
				"error": apperrors.ErrMissingFields.Error(),
				"details": map[string]interface{}{
					"Age":  "required, int",
					"Date": "required, time.Time",
				},
			}

			utils.FormatAndSendRequiredFieldsError(err, ctx)

			assert.Equal(tt, http.StatusBadRequest, w.Code, "Status code should be 400 (bad request)")

			body := map[string]interface{}{}
			json.Unmarshal(w.Body.Bytes(), &body)

			assert.Equal(tt, expectedBody, body, "expectedBody should match the writter body")
		},
	)

	t.Run("Should return a invalid request error when fail on validate struct",
		func(tt *testing.T) {
			setUp()

			err := errors.New("random error")

			expectedBody := map[string]interface{}{
				"error":   apperrors.ErrInvalidRequest.Error(),
				"details": err.Error(),
			}

			utils.FormatAndSendRequiredFieldsError(err, ctx)

			assert.Equal(tt, http.StatusBadRequest, w.Code, "Status code should be 400 (bad request)")

			body := map[string]interface{}{}
			json.Unmarshal(w.Body.Bytes(), &body)

			assert.Equal(tt, expectedBody, body, "expectedBody should match the writter body")
		},
	)

}
