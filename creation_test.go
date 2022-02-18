package sterrors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
	err := fmt.Errorf("error cause")
	severity := SeverityError
	const testKind Kind = 0

	resultErr := E(message, err, severity, testKind)

	assert.IsType(t, &BaseError{}, resultErr)
	e := resultErr.(*BaseError)

	assert.Equal(t, e.Message, message)
	assert.Equal(t, e.Cause, err)
	assert.Equal(t, e.Severity, severity)
	assert.Equal(t, e.Kind, testKind)
}

func TestEWithCustomErrorType(t *testing.T) {
	SetDefaultCreateErrorFunc(func() Error {
		return &CustomErrorType{}
	})

	testCases := map[string]struct {
		Message         string
		Severity        Severity
		Kind            Kind
		CustomAttribute CustomAttribute
	}{
		"Simple custom error type": {
			Message:         "Error message",
			Severity:        SeverityError,
			Kind:            0,
			CustomAttribute: 1,
		},
		"Cause is also a custom error": {
			Message:         "Error message",
			Severity:        SeverityError,
			Kind:            0,
			CustomAttribute: 1,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			cause := E("cause message", SeverityWarning)
			causeCaller := func() Caller { return caller() }()
			causeCaller.Line--
			expectedCauseCallStackEntry := CallStackEntry{
				ErrMessage: "cause message",
				Caller:     causeCaller,
			}
			resultErr := E(tc.Message, cause, tc.Severity, tc.Kind, tc.CustomAttribute)
			errCaller := func() Caller { return caller() }()
			errCaller.Line--
			expectedCallStackEntry := CallStackEntry{
				ErrMessage: tc.Message,
				Caller:     errCaller,
			}

			assert.IsType(t, &CustomErrorType{}, resultErr)
			e := resultErr.(*CustomErrorType)

			callStack := e.CallStack()
			assert.Len(t, callStack, 2)
			assert.Equal(t, expectedCallStackEntry, callStack[0])
			assert.Equal(t, expectedCauseCallStackEntry, callStack[1])

			assert.Equal(t, e.Message, tc.Message)
			assert.Equal(t, e.Cause, cause)
			assert.Equal(t, e.Severity, tc.Severity)
			assert.Equal(t, e.Kind, tc.Kind)
			assert.Equal(t, e.CustomAttribute, tc.CustomAttribute)
		})
	}
}
