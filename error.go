package sterrors

type Error interface {
	error
	CallStack() []CallStackEntry
	Cause() error
	Severity() Severity
	Kind() Kind
	Caller() Caller
	enrich(args ...interface{})
	setCaller(caller Caller)
}

type BaseError struct {
	message  string
	caller   Caller
	kind     Kind
	severity Severity
	cause    error
}

type Caller struct {
	FuncName string `json:"funcName,omitempty"`
	File     string `json:"file,omitempty"`
	Line     int    `json:"line,omitempty"`
}

type CallStackEntry struct {
	ErrMessage string `json:"msg,omitempty"`
	Caller     Caller `json:"caller,omitempty"`
}

type Kind string
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

func (e *BaseError) Kind() Kind {
	return e.kind
}

func (e *BaseError) Caller() Caller {
	return e.caller
}

// CallStack returns the callstack
// It adds CallStack entries recursively based on error causes
// as long they also implement the Error interface
func (e *BaseError) CallStack() []CallStackEntry {
	res := []CallStackEntry{{ErrMessage: e.Error(), Caller: e.caller}}

	if e.cause == nil {
		return res
	}

	subErr, ok := e.cause.(Error)
	if !ok {
		res = append(res, CallStackEntry{ErrMessage: e.cause.Error()})
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
	e.BaseError.enrich(args...)
}
*/
func (e *BaseError) enrich(args ...interface{}) {
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
