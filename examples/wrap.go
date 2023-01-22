package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Bibob7/sterrors"
)

var WrappingError = errors.New("wrapping error")

func main() {
	err := anotherMethod()

	// second error that results from the first one
	second := sterrors.Wrap(WrappingError, "action not possible", err)

	jsonStackTrace, _ := json.Marshal(sterrors.CallStack(second))
	// Print out the error stack trace
	fmt.Printf("%s \n", jsonStackTrace)
	// Calling Error() return the wrapped error message
	fmt.Printf("Simply print the wrapped error message: %s \n", second.Error())
	// The provided error type also implements the Unwrap() error method so that you can use it with errors.Is()
	fmt.Printf("Is wrapping error: %v \n", errors.Is(second, WrappingError))
}

func anotherMethod() error {
	return sterrors.E("some error message")
}

/*
Output:
[
   {
      "msg":"wrapping error",
      "caller":{
         "funcName":"main.main",
         "file":"/Users/kevin/Repositories/sterrors/examples/wrap.go",
         "line":16
      }
   },
   {
      "msg":"action not possible",
      "caller":{
         "funcName":"main.main",
         "file":"/Users/kevin/Repositories/sterrors/examples/wrap.go",
         "line":16
      }
   },
   {
      "msg":"some error message",
      "caller":{
         "funcName":"main.anotherMethod",
         "file":"/Users/kevin/Repositories/sterrors/examples/wrap.go",
         "line":28
      }
   }
]

Simply print the wrapped error message: wrapping error: action not possible: some error message

Is wrapping error: true
*/
