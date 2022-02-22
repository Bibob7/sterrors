package sterrors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func callFake() Caller {
	return caller()
}

func TestCallStack(t *testing.T) {
	SetDefaultCreateErrorFunc(func() Error {
		return &BaseError{}
	})
	initErr := fmt.Errorf("initial error")
	secErr, secondCall := E("second error", initErr), callFake()
	thirdErr, thirdCall := E("third error", secErr), callFake()

	callStack := CallStack(thirdErr)

	assert.Len(t, callStack, 3)
	assert.Equal(t, "third error", callStack[0].ErrMessage)
	assert.Equal(t, "second error", callStack[1].ErrMessage)
	assert.Equal(t, "initial error", callStack[2].ErrMessage)

	assert.Equal(t, thirdCall, callStack[0].Caller)
	assert.Equal(t, secondCall, callStack[1].Caller)
	assert.Equal(t, Caller{}, callStack[2].Caller)
}

func TestCallStack_WithNotTraceableErr(t *testing.T) {
	err := fmt.Errorf("initial error")

	callStack := CallStack(err)

	assert.Len(t, callStack, 1)
	assert.Equal(t, "initial error", callStack[0].ErrMessage)
}

func TestCallStack_WithErrWithoutCause(t *testing.T) {
	err := E("initial error")

	callStack := CallStack(err)

	assert.Len(t, callStack, 1)
	assert.Equal(t, "initial error", callStack[0].ErrMessage)
}

func TestCallStack_WithNil(t *testing.T) {
	callStack := CallStack(nil)

	assert.Len(t, callStack, 0)
}
