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
