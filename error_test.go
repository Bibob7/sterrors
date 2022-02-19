package sterrors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func callFake() Caller {
	return caller()
}

func TestBaseError_CallStack(t *testing.T) {
	initErr := fmt.Errorf("initial error")
	secErr, secondCall := E("second error", initErr), callFake()
	thirdErr, thirdCall := E("third error", secErr), callFake()

	e, ok := thirdErr.(Error)
	assert.True(t, ok)

	callStack := e.CallStack()

	assert.Len(t, callStack, 3)
	assert.Equal(t, callStack[0].ErrMessage, "third error")
	assert.Equal(t, callStack[1].ErrMessage, "second error")
	assert.Equal(t, callStack[2].ErrMessage, "initial error")

	assert.Equal(t, callStack[0].Caller, thirdCall)
	assert.Equal(t, callStack[1].Caller, secondCall)
	assert.Equal(t, callStack[2].Caller, Caller{})
}

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
