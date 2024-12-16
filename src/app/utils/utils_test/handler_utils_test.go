package utils_test

import (
	"errors"
	"reflect"
	"testing"

	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/models"
	"sm.com/m/src/app/utils"
)

type InputGetRequiredFields struct {
	RequiredFields []string
	Data           interface{}
}

func TestGetRequiredFields(t *testing.T) {
	t.Run("Should return one missing field", func(t *testing.T) {
		input := InputGetRequiredFields{
			RequiredFields: []string{"Name", "Email", "Password"},
			Data:           models.UserSignIn{Name: "josh", Password: "123"},
		}
		want := map[string]any{"email": "required"}

		result, _ := utils.GetMissingFields(input.RequiredFields, input.Data)

		if !reflect.DeepEqual(result, want) {
			t.Errorf("ERROR: \nINPUT: %v\nWANT: %v\nGOT: %v\n", input, want, result)
		}
	})

	t.Run("Should return two missing fields", func(t *testing.T) {
		input := InputGetRequiredFields{
			RequiredFields: []string{"Name", "Email", "Password"},
			Data:           models.UserSignIn{Name: "josh"},
		}
		want := map[string]any{"email": "required", "password": "required"}

		result, _ := utils.GetMissingFields(input.RequiredFields, input.Data)

		if !reflect.DeepEqual(result, want) {
			t.Errorf("ERROR: \nINPUT: %v\nWANT: %v\nGOT: %v\n", input, want, result)
		}
	})

	t.Run("Should return all fields as missing", func(t *testing.T) {
		input := InputGetRequiredFields{
			RequiredFields: []string{"Name", "Email", "Password"},
			Data:           models.UserSignIn{},
		}
		want := map[string]any{"name": "required", "email": "required", "password": "required"}

		result, _ := utils.GetMissingFields(input.RequiredFields, input.Data)

		if !reflect.DeepEqual(result, want) {
			t.Errorf("ERROR: \nINPUT: %v\nWANT: %v\nGOT: %v\n", input, want, result)
		}
	})

	t.Run("Should not return missing fields", func(t *testing.T) {
		input := InputGetRequiredFields{
			RequiredFields: []string{"Name", "Email", "Password"},
			Data:           models.UserSignIn{Name: "josh", Email: "josh@gmail.com", Password: "123"},
		}
		want := map[string]any{}

		result, _ := utils.GetMissingFields(input.RequiredFields, input.Data)

		if !reflect.DeepEqual(result, want) {
			t.Errorf("ERROR: \nINPUT: %v\nWANT: %v\nGOT: %v\n", input, want, result)
		}
	})

	t.Run("Should return a error: data is not a struct", func(t *testing.T) {
		input := InputGetRequiredFields{
			RequiredFields: []string{"Name", "Email", "Password"},
			Data:           "not a struct",
		}
		want := apperrors.ErrIsNotStruct

		_, err := utils.GetMissingFields(input.RequiredFields, input.Data)
		if err == nil || !errors.Is(want, err) {
			t.Errorf("ERROR: \nINPUT: %v\nWANT: %v\nGOT: %v\n", input, want, err)
		}
	})
}
