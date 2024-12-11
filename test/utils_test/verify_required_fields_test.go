package utils_test

import (
	"reflect"
	"testing"

	"sm.com/m/src/app/utils"
)

func TestVerifyRequiredFields(t *testing.T) {
	type TestInput struct {
		name           string
		requiredFields []string
		input          map[string]any
		expected       map[string]string
	}

	tests := []TestInput{{
		name:           "Misssing one field",
		requiredFields: []string{"name", "age", "gender"},
		input:          map[string]any{"name": "josh", "age": 10},
		expected:       map[string]string{"gender": "required"},
	}, {
		name:           "Misssing two fields",
		requiredFields: []string{"name", "age", "gender"},
		input:          map[string]any{"name": "josh"},
		expected:       map[string]string{"age": "required", "gender": "required"},
	}, {
		name:           "Misssing all fields",
		requiredFields: []string{"name", "age", "gender"},
		input:          map[string]any{},
		expected:       map[string]string{"name": "required", "age": "required", "gender": "required"},
	}, {
		name:           "Not missing fields",
		requiredFields: []string{},
		input:          map[string]any{"name": "josh", "age": 10, "gender": "male"},
		expected:       map[string]string{},
	},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := utils.VerifyRequiredFields(test.requiredFields, test.input)
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("VerifyRequiredFields(%q, %q) = %q; want %q;", test.requiredFields, test.input, result, test.expected)
			}
		})
	}

}
