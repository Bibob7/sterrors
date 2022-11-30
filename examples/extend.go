package main

import (
	"fmt"
	"github.com/Bibob7/sterrors"
)

type MyError struct {
	sterrors.BaseError
	CustomAttribute CustomAttribute
}

type CustomAttribute int

func (e *MyError) Enrich(args ...interface{}) {
	e.BaseError.Enrich()
	for _, arg := range args {
		switch arg := arg.(type) {
		case CustomAttribute:
			e.CustomAttribute = arg
		}
	}
}

func main() {
	sterrors.SetDefaultCreateErrorFunc(func() sterrors.Error {
		return &MyError{}
	})
	var custom CustomAttribute = 2

	err := sterrors.E("some message", custom)

	e, _ := err.(*MyError)

	fmt.Printf("%d \n", e.CustomAttribute)
}
