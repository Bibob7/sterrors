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
	e.setCaller(caller())
	e.enrich(args...)
	return e
}
