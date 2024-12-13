package utils_test

import (
	"reflect"
	"testing"

	"sm.com/m/src/app/utils"
)

func TestGetRequiredFields(t *testing.T) {
	type InputStruct struct {
		Name     string
		Email    string
		Password string
	}

	t.Run("Missing one field", func(t *testing.T) {
		requiredFields := []string{"Name", "Email", "Password"}
		input := &InputStruct{Name: "josh", Password: "123"}
		expected := map[string]any{"email": "required"}

		result := utils.GetMissingFields(requiredFields, input)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("VerifyRequiredFields(%q, %q) = %q; want %q;", requiredFields, input, result, expected)
		}
	})

	t.Run("Missing two fields", func(t *testing.T) {
		requiredFields := []string{"Name", "Email", "Password"}
		input := &InputStruct{Name: "josh"}
		expected := map[string]any{"email": "required", "password": "required"}

		result := utils.GetMissingFields(requiredFields, input)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("VerifyRequiredFields(%q, %q) = %q; want %q;", requiredFields, input, result, expected)
		}
	})

	t.Run("Missing all fields", func(t *testing.T) {
		requiredFields := []string{"Name", "Email", "Password"}
		input := &InputStruct{}
		expected := map[string]any{"name": "required", "email": "required", "password": "required"}

		result := utils.GetMissingFields(requiredFields, input)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("VerifyRequiredFields(%q, %q) = %q; want %q;", requiredFields, input, result, expected)
		}
	})

	t.Run("Not missing fields", func(t *testing.T) {
		requiredFields := []string{"Name", "Email", "Password"}
		input := &InputStruct{Name: "josh", Email: "josh@gmail.com", Password: "123"}
		expected := map[string]any{}

		result := utils.GetMissingFields(requiredFields, input)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("VerifyRequiredFields(%q, %q) = %q; want %q;", requiredFields, input, result, expected)
		}
	})
}
