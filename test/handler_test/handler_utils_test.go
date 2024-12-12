package utils_test

import (
	"reflect"
	"testing"

	"sm.com/m/src/app/handlers"
)

func TestGetRequiredFields(t *testing.T) {
	t.Run("Missing one field", func(t *testing.T) {
		requiredFields := []string{"name", "email", "phone"}
		input := map[string]any{"name": "", "email": ""}
		expected := map[string]any{"phone": "required"}

		result := handlers.GetRequiredFields(requiredFields, input)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("VerifyRequiredFields(%q, %q) = %q; want %q;", requiredFields, input, result, expected)
		}
	})

	t.Run("Missing two fields", func(t *testing.T) {
		requiredFields := []string{"name", "email", "phone"}
		input := map[string]any{"name": ""}
		expected := map[string]any{"email": "required", "phone": "required"}

		result := handlers.GetRequiredFields(requiredFields, input)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("VerifyRequiredFields(%q, %q) = %q; want %q;", requiredFields, input, result, expected)
		}
	})

	t.Run("Missing all fields", func(t *testing.T) {
		requiredFields := []string{"name", "email", "phone"}
		input := map[string]any{}
		expected := map[string]any{"name": "required", "email": "required", "phone": "required"}

		result := handlers.GetRequiredFields(requiredFields, input)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("VerifyRequiredFields(%q, %q) = %q; want %q;", requiredFields, input, result, expected)
		}
	})

	t.Run("Not missing fields", func(t *testing.T) {
		requiredFields := []string{"name", "email", "phone"}
		input := map[string]any{"name": "", "email": "", "phone": ""}
		expected := map[string]any{}

		result := handlers.GetRequiredFields(requiredFields, input)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("VerifyRequiredFields(%q, %q) = %q; want %q;", requiredFields, input, result, expected)
		}
	})
}
