package sterrors

type Error interface {
	error
	Cause() error
	Severity() Severity
	Kind() Kind
	Caller() *Caller
	Enrich(args ...interface{})
	setCaller(caller *Caller)
}

type BaseError struct {
	message  string
	kind     Kind
	caller   *Caller
	severity Severity
	cause    error
}

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
		case Severity:
			e.severity = arg
		case Kind:
			e.kind = arg
		case string:
			e.message = arg
		default:
			// ignore unknown arg types
		}
	}
}

// setCaller is used during the creation to set the caller
func (e *BaseError) setCaller(caller *Caller) {
	e.caller = caller
}
