package sterrors

var (
	createError CreateFunc = func() Error {
		return &BaseError{}
	}
)

type CreateFunc func() Error

// SetDefaultCreateErrorFunc sets the default error type
func SetDefaultCreateErrorFunc(customCreateErrorFunc CreateFunc) {
	createError = customCreateErrorFunc
}

// E is used to create a new error
// All potential error attributes can be passed in random order
func E(args ...interface{}) error {
	e := createError()
	e.setCaller(caller(2))
	e.Enrich(args...)
	return e
}

// Wrap is used to create a new error based on an existing error, which has to be passed as first argument
// All potential error attributes can be passed in random order
func Wrap(err error, args ...interface{}) error {
	e := createError()
	e.setCaller(caller(2))
	e.Wrap(err, args...)
	return e
}
