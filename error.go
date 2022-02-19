package sterrors

type Error interface {
	error
	CallStack() []CallStackEntry
	Enrich(args ...interface{})
	setCaller(caller Caller)
	Cause() error
	Severity() Severity
}

type BaseError struct {
	message  string
	caller   Caller
	kind     Kind
	severity Severity
	cause    error
}

type Caller struct {
	FuncName string
	File     string
	Line     int
}

type CallStackEntry struct {
	ErrMessage string
	Caller     Caller
}

type Kind int
type Severity int

func (e *BaseError) Error() string {
	return e.message
}

func (e *BaseError) Cause() error {
	return e.cause
}

func (e *BaseError) Severity() Severity {
	return e.severity
}

// CallStack returns the callstack
// It adds CallStack entries recursively based on error causes
// as long they also implement the Error interface
func (e *BaseError) CallStack() []CallStackEntry {
	res := []CallStackEntry{{ErrMessage: e.Error(), Caller: e.caller}}

	subErr, ok := e.cause.(Error)
	if !ok {
		if e.cause != nil {
			res = append(res, CallStackEntry{ErrMessage: e.cause.Error()})
		}
		return res
	}

	res = append(res, subErr.CallStack()...)

	return res
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
		case Kind:
			e.kind = arg
		case Severity:
			e.severity = arg
		case string:
			e.message = arg
		default:
			// ignore unknown arg types
		}
	}
}

// setCaller is used during the creation to set the caller
func (e *BaseError) setCaller(caller Caller) {
	e.caller = caller
}
