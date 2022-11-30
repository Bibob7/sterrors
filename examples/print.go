package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Bibob7/sterrors"
)

var InitialError = errors.New("initial error")

func main() {
	err := secondMethod()

	// second error that results from the first one
	second := sterrors.E("action not possible", err)

	jsonStackTrace, _ := json.Marshal(sterrors.CallStack(second))
	// Print out the error stack trace
	fmt.Printf("%s \n", jsonStackTrace)
	// Calling Error() return the wrapped error message
	fmt.Printf("Simply print the wrapped error message: %s \n", second.Error())
	// The provided error type also implements the Unwrap() error method so that you can use it with errors.Is()
	fmt.Printf("Is initial error: %v \n", errors.Is(second, InitialError))
}

func secondMethod() error {
	return sterrors.E("some error message", InitialError)
}

/*
Output:
[
  {
    "msg": "action not possible",
    "caller": {
      "funcName": "main.main",
      "file": "/Users/root/Repositories/sterrors/examples/main.go",
      "line": 13
    }
  },
  {
    "msg": "some error message",
    "caller": {
      "funcName": "main.anotherMethod",
      "file": "/Users/root/Repositories/sterrors/examples/main.go",
      "line": 20
    }
  }
]

Simply print the wrapped error message: action not possible: some error message

Is initial error: true
*/
