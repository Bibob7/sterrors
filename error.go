package sterrors

import "fmt"

type Error interface {
	error
	Message() string
	Caller() *Caller
	Unwrap() error
	Enrich(args ...interface{})
	Wrap(err error, args ...interface{})
	Is(err error) bool
	setCaller(caller *Caller)
}

type BaseError struct {
	message string
	caller  *Caller
	cause   error
}

func (e *BaseError) Message() string {
	return e.message
}

func (e *BaseError) Error() string {
	if e.cause != nil {
		if e.message == "" {
			return e.cause.Error()
		}
		return fmt.Sprintf("%s: %s", e.Message(), e.cause.Error())
	}
	return e.Message()
}

func (e *BaseError) Unwrap() error {
	return e.cause
}

func (e *BaseError) Caller() *Caller {
	return e.caller
}

/*
Enrich can be overridden or extended by a custom error type

For example:

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
*/
func (e *BaseError) Enrich(args ...interface{}) {
	for _, arg := range args {
		switch arg := arg.(type) {
		case error:
			e.cause = arg
		case string:
			e.message = arg
		default:
			// ignore unknown arg types
		}
	}
}
func (e *BaseError) Is(err error) bool {
	return e.message == err.Error()
}

func (e *BaseError) Wrap(err error, args ...interface{}) {
	if err == nil {
		return
	}
	numArgs := len(args)
	e.message = err.Error()
	if numArgs == 0 {
		return
	}
	if numArgs == 1 {
		if cause, ok := args[0].(error); ok {
			e.cause = cause
			return
		}
	}
	wrappedErr := createError()
	wrappedErr.setCaller(caller(3))
	wrappedErr.Enrich(args...)
	e.cause = wrappedErr
}

// setCaller is used during the creation to set the caller
func (e *BaseError) setCaller(caller *Caller) {
	e.caller = caller
}
