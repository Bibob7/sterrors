package sterrors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBaseError_Enrich(t *testing.T) {
	message := "error message"
	cause := fmt.Errorf("initial error")
	severity := SeverityInfo

	e := BaseError{}
	e.Enrich(message, cause, severity)

	assert.Equal(t, e.message, message)
	assert.Equal(t, e.severity, severity)
	assert.Equal(t, e.cause, cause)
}

func TestBaseError_Error(t *testing.T) {
	testCases := map[string]struct {
		Message              string
		ExpectedErrorMessage string
	}{
		"Custom error message": {
			Message:              "custom error message",
			ExpectedErrorMessage: "custom error message",
		},
		"No custom error message": {
			ExpectedErrorMessage: "unexpected",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := E(tc.Message)
			assert.Equal(t, tc.ExpectedErrorMessage, err.Error())
		})
	}
}
