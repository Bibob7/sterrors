package sterrors

import (
	"fmt"
	"reflect"
	"testing"
)

type CustomErrorType struct {
	BaseError
	CustomAttribute
}

type CustomAttribute int

func (e *CustomErrorType) Enrich(args ...interface{}) {
	for _, arg := range args {
		switch arg := arg.(type) {
		case CustomAttribute:
			e.CustomAttribute = arg
		}
	}
	e.BaseError.Enrich(args...)
}

func TestE(t *testing.T) {
	message := "Error message"
	cause := fmt.Errorf("error Cause")
	severity := SeverityError

	resultErr := E(message, cause, severity)

	if reflect.TypeOf(resultErr) != reflect.TypeOf(&BaseError{}) {
		t.Errorf("error has no type BaseError")
	}
	err := resultErr.(*BaseError)

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

func TestEWithCustomErrorType(t *testing.T) {
	SetDefaultCreateErrorFunc(func() Error {
		return &CustomErrorType{}
	})

	testCases := map[string]struct {
		Message         string
		Severity        Severity
		CustomAttribute CustomAttribute
	}{
		"Simple custom error type": {
			Message:         "Error message",
			Severity:        SeverityError,
			CustomAttribute: 1,
		},
		"Cause is also a custom error": {
			Message:         "Error message",
			Severity:        SeverityError,
			CustomAttribute: 1,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			cause := E("Cause message", SeverityWarning)
			resultErr := E(tc.Message, cause, tc.Severity, tc.CustomAttribute)

			if reflect.TypeOf(resultErr) != reflect.TypeOf(&CustomErrorType{}) {
				t.Errorf("error has no type CustomErrorType")
			}
			e := resultErr.(*CustomErrorType)

			if e.message != tc.Message {
				t.Errorf("error message is not equal")
			}
			if e.cause != cause {
				t.Errorf("error cause is not cause")
			}
			if e.severity != tc.Severity {
				t.Errorf("error cause is not cause")
			}
			if e.CustomAttribute != tc.CustomAttribute {
				t.Errorf("error cause is not cause")
			}
		})
	}
}
