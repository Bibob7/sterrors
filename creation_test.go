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

	resultErr := E(message, cause)

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
}

func TestEWithCustomErrorType(t *testing.T) {
	SetDefaultCreateErrorFunc(func() Error {
		return &CustomErrorType{}
	})

	errMsg := "Error message"
	customAttribute := CustomAttribute(1)
	cause := E("Cause message")
	resultErr := E(errMsg, cause, customAttribute)

	if reflect.TypeOf(resultErr) != reflect.TypeOf(&CustomErrorType{}) {
		t.Errorf("error has no type CustomErrorType")
	}
	e := resultErr.(*CustomErrorType)

	if e.message != errMsg {
		t.Errorf("error message is not equal")
	}
	if e.cause != cause {
		t.Errorf("error cause is not cause")
	}
	if e.CustomAttribute != customAttribute {
		t.Errorf("error custom attribute is not custom attribute")
	}
}
