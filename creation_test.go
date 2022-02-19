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

func (e *CustomErrorType) enrich(args ...interface{}) {
	for _, arg := range args {
		switch arg := arg.(type) {
		case CustomAttribute:
			e.CustomAttribute = arg
		}
	}
	e.BaseError.enrich(args...)
}

func TestE(t *testing.T) {
	message := "Error message"
	err := fmt.Errorf("error Cause")
	severity := SeverityError
	const testKind Kind = 0

	resultErr := E(message, err, severity, testKind)

	assert.IsType(t, &BaseError{}, resultErr)
	e := resultErr.(*BaseError)

	assert.Equal(t, e.message, message)
	assert.Equal(t, e.cause, err)
	assert.Equal(t, e.severity, severity)
	assert.Equal(t, e.kind, testKind)
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
			cause := E("Cause message", SeverityWarning)
			resultErr := E(tc.Message, cause, tc.Severity, tc.Kind, tc.CustomAttribute)

			assert.IsType(t, &CustomErrorType{}, resultErr)
			e := resultErr.(*CustomErrorType)

			assert.Equal(t, e.message, tc.Message)
			assert.Equal(t, e.cause, cause)
			assert.Equal(t, e.severity, tc.Severity)
			assert.Equal(t, e.kind, tc.Kind)
			assert.Equal(t, e.CustomAttribute, tc.CustomAttribute)
		})
	}
}
