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
			ExpectedErrorMessage: "unexpected",
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
			if tc.Message == "" && stError.Message() != "unexpected" {
				t.Errorf("empty message is not mapped to unexpected")
			}
		})
	}
}
