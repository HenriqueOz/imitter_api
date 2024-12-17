package utils_test

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"reflect"
	"testing"

	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/models"
	"sm.com/m/src/app/utils"
)

type InputSendSuccess struct {
	recorder *httptest.ResponseRecorder
	payload  interface{}
	status   int
}

func TestSendSuccess(t *testing.T) {
	input := InputSendSuccess{
		recorder: httptest.NewRecorder(),
		payload: models.UserSignUp{
			Name: "Josh", Email: "josh@gmail.com", Password: "123",
		},
		status: 200,
	}
	var want error = nil

	err := utils.SendSuccess(input.recorder, input.payload, input.status)

	if err != nil {
		t.Errorf("ERROR: should not return an error\nINPUT: %v\nWANT: %v\nGOT: %v\n", input, want, err)
	}

	recorderStatusCode := input.recorder.Result().StatusCode
	if recorderStatusCode != input.status {
		t.Errorf("ERROR: status code is different \nINPUT: %v\nWANT: %v\nGOT: %v\n", input, input.status, recorderStatusCode)
	}

	var decodedBody models.UserSignUp
	json.Unmarshal(input.recorder.Body.Bytes(), &decodedBody)
	if !reflect.DeepEqual(decodedBody, input.payload) {
		t.Errorf("ERROR: output payload is different \nINPUT: %v\nWANT: %v\nGOT: %v\n", input, input.payload, decodedBody)
	}
}

type InputGetRequiredFields struct {
	RequiredFields []string
	Data           interface{}
}

func TestGetRequiredFields(t *testing.T) {
	t.Run("Should return one missing field", func(t *testing.T) {
		input := InputGetRequiredFields{
			RequiredFields: []string{"Name", "Email", "Password"},
			Data:           models.UserSignUp{Name: "josh", Password: "123"},
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
			Data:           models.UserSignUp{Name: "josh"},
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
			Data:           models.UserSignUp{},
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
			Data:           models.UserSignUp{Name: "josh", Email: "josh@gmail.com", Password: "123"},
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
