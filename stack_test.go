package sterrors

import (
	"fmt"
	"reflect"
	"testing"
)

func callFake() *Caller {
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

	if len(callStack) != 3 {
		t.Errorf("callstack has not lenth 3")
	}
	if "third error" != callStack[0].ErrMessage {
		t.Errorf("third callstack error message is not equal to third error")
	}
	if "second error" != callStack[1].ErrMessage {
		t.Errorf("second callstack error message is not equal to second error")
	}
	if "initial error" != callStack[2].ErrMessage {
		t.Errorf("first callstack error message is not equal to initial error")
	}

	var nilCaller *Caller
	if !reflect.DeepEqual(thirdCall, callStack[0].Caller) {
		t.Errorf("third caller is not equal to actual first caller")
	}
	if !reflect.DeepEqual(secondCall, callStack[1].Caller) {
		t.Errorf("second caller is not equal to actual second caller")
	}
	if !reflect.DeepEqual(nilCaller, callStack[2].Caller) {
		t.Errorf("first caller is not equal to nilCaller")
	}
}

func TestCallStack_WithNotTraceableErr(t *testing.T) {
	err := fmt.Errorf("initial error")

	callStack := CallStack(err)

	if len(callStack) != 1 {
		t.Errorf("callstack has not lenth 1")
	}
	if "initial error" != callStack[0].ErrMessage {
		t.Errorf("first callstack error message is not equal to initial error")
	}
}

func TestCallStack_WithErrWithoutCause(t *testing.T) {
	err := E("initial error")

	callStack := CallStack(err)

	if len(callStack) != 1 {
		t.Errorf("callstack has not lenth 1")
	}
	if "initial error" != callStack[0].ErrMessage {
		t.Errorf("first callstack error message is not equal to initial error")
	}
}

func TestCallStack_WithNil(t *testing.T) {
	callStack := CallStack(nil)

	if len(callStack) != 0 {
		t.Errorf("callstack has not lenth 0")
	}
}
