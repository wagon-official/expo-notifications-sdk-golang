package expo

import (
	"encoding/json"
	"strings"
	"testing"
)

type Example struct {
	Token ExpoPushToken `json:"token"`
}

func TestUnmarshallingExponentPushNotification(t *testing.T) {
	jsonString := `{ "token" : "ExponentPushToken[xxxxxxxxxxxxxxxxxxxxxx]" }`
	reader := strings.NewReader(jsonString)
	var example Example
	err := json.NewDecoder(reader).Decode(&example)
	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshallingFailExponentPushNotification(t *testing.T) {
	jsonString := `{ "token" : "BadToken" }`
	reader := strings.NewReader(jsonString)
	var example Example
	err := json.NewDecoder(reader).Decode(&example)
	if err != ErrMalformedToken {
		t.FailNow()
	}
}

func TestValidateResponseErrorStatus(t *testing.T) {
	response := &PushResponse{
		Status:  "error",
		Message: "failed",
		Details: map[string]interface{}{},
	}
	err := response.ValidateResponse()
	typed, ok := err.(*PushResponseError)
	if !ok {
		t.Error("Incorrect error type")
	}
	if typed.Response != response {
		t.Error("Didn't return called response")
	}
}

func TestValidateResponseSuccess(t *testing.T) {
	response := &PushResponse{
		Status: "ok",
	}
	err := response.ValidateResponse()
	if err != nil {
		t.Error("Errored on valid response")
	}
}

func TestValidateResponseDeviceNotRegistered(t *testing.T) {
	response := &PushResponse{
		Status:  "error",
		Message: "Not registered",
		Details: map[string]interface{}{"error": "DeviceNotRegistered"},
	}
	err := response.ValidateResponse()
	typed, ok := err.(*DeviceNotRegisteredError)
	if !ok {
		t.Error("Incorrect error type")
	}
	if typed.Response != response {
		t.Error("Didn't return called response")
	}
}

func TestValidateResponseErrorMessageTooBig(t *testing.T) {
	response := &PushResponse{
		Status:  "error",
		Message: "Message too big",
		Details: map[string]interface{}{"error": "MessageTooBig"},
	}
	err := response.ValidateResponse()
	typed, ok := err.(*MessageTooBigError)
	if !ok {
		t.Error("Incorrect error type")
	}
	if typed.Response != response {
		t.Error("Didn't return called response")
	}
}

func TestValidateResponseErrorMessageRateExceeded(t *testing.T) {
	response := &PushResponse{
		Status:  "error",
		Message: "Too many messages at once",
		Details: map[string]interface{}{"error": "MessageRateExceeded"},
	}
	err := response.ValidateResponse()
	typed, ok := err.(*MessageRateExceededError)
	if !ok {
		t.Error("Incorrect error type")
	}
	if typed.Response != response {
		t.Error("Didn't return called response")
	}
}
