package sterrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestBaseError_Enrich(t *testing.T) {
	message := "error message"
	cause := fmt.Errorf("initial error")

	err := BaseError{}
	err.Enrich(message, cause)

	if err.message != message {
		t.Errorf("error message is not equal")
	}
	if err.cause != cause {
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

	testCases := map[string]struct {
		Error            error
		Message          string
		CallStackEntries int
		IsCustomError    bool
		IsAnotherError   bool
	}{
		"With custom error message": {
			Error:            Wrap(CustomError, "some additional error message", AnotherError),
			Message:          "custom error: some additional error message: another error",
			CallStackEntries: 3,
			IsCustomError:    true,
			IsAnotherError:   true,
		},
		"Wrap with previous error": {
			Error:            Wrap(CustomError, AnotherError),
			Message:          "custom error: another error",
			CallStackEntries: 2,
			IsCustomError:    true,
			IsAnotherError:   true,
		},
		"Wrap only": {
			Error:            Wrap(CustomError),
			Message:          "custom error",
			CallStackEntries: 1,
			IsCustomError:    true,
			IsAnotherError:   false,
		},
		"Wrap nothing": {
			Error:            Wrap(nil),
			Message:          "",
			CallStackEntries: 1,
			IsCustomError:    false,
			IsAnotherError:   false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if tc.IsCustomError && !errors.Is(tc.Error, CustomError) {
				t.Errorf("error is not custom error")
			}
			if tc.IsAnotherError && !errors.Is(tc.Error, AnotherError) {
				t.Errorf("error is not another error")
			}
			if tc.Error.Error() != tc.Message {
				t.Errorf("error message is \"%s\" instead of \"%s\"", tc.Error.Error(), tc.Message)
			}
			callStack := CallStack(tc.Error)
			numberOfEntries := len(callStack)
			if numberOfEntries != tc.CallStackEntries {
				t.Errorf("number of callstack entries is \"%d\" not 3", numberOfEntries)
			}
		})
	}
}
