package sterrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestBaseError_Enrich(t *testing.T) {
	message := "error message"
	cause := fmt.Errorf("initial error")
	severity := SeverityInfo

	err := BaseError{}
	err.Enrich(message, cause, severity)

	if err.message != message {
		t.Errorf("error message is not equal")
	}
	if err.cause != cause {
		t.Errorf("error cause is not cause")
	}
	if err.severity != severity {
		t.Errorf("error cause is not cause")
	}
}

func TestBaseError_Error(t *testing.T) {
	testCases := map[string]struct {
		Message              string
		ExpectedErrorMessage string
		Cause                error
	}{
		"Custom error message": {
			Message:              "custom error message",
			ExpectedErrorMessage: "custom error message",
		},
		"No custom error message": {
			ExpectedErrorMessage: "",
		},
		"With wrapped error": {
			Message:              "some message",
			ExpectedErrorMessage: "some message: previous error",
			Cause:                errors.New("previous error"),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := E(tc.Message)
			if tc.Cause != nil {
				err = E(tc.Message, tc.Cause)
			}
			if tc.ExpectedErrorMessage != err.Error() {
				t.Errorf("expected message \"%s\" is not actual message \"%s\"", tc.ExpectedErrorMessage, err.Error())
			}
			stError, _ := err.(Error)
			if tc.Message != "" && tc.Message != stError.Message() {
				t.Errorf("expected message \"%s\" is not actual message \"%s\"", tc.Message, stError.Message())
			}
		})
	}
}

func TestError_Wrap(t *testing.T) {
	AnotherError := errors.New("another error")
	CustomError := errors.New("custom error")
	err := Wrap(CustomError, AnotherError)

	if !errors.Is(err, CustomError) {
		t.Errorf("error is not custom error")
	}
	if !errors.Is(err, AnotherError) {
		t.Errorf("error is not another error")
	}
	if err.Error() != "custom error: another error" {
		t.Errorf("error message is \"%s\" instead of \"%s\"", err.Error(), "custom error: another error")
	}
}
