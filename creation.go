package sterrors

import (
	"runtime"
	"strings"
)

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
	e.Enrich(args...)
	return e
}

func caller() Caller {
	pc, file, line, _ := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	nameSegments := strings.Split(details.Name(), "/")
	funcName := details.Name()
	if len(nameSegments) > 0 {
		funcName = nameSegments[len(nameSegments)-1]
	}
	return Caller{File: file, Line: line, FuncName: funcName}
}
