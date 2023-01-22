# Sterrors

[![codecov](https://codecov.io/gh/Bibob7/sterrors/branch/main/graph/badge.svg?token=2LURD0VD9X)](https://codecov.io/gh/Bibob7/sterrors)
[![CircleCI](https://circleci.com/gh/Bibob7/sterrors/tree/main.svg?style=svg)](https://circleci.com/gh/Bibob7/sterrors/tree/main)

sterrors is a simple library which will provide error types with stack traces and severity levels.

You can extend the provides BaseError type and add you own application specific attributes that you may want to log.

## Error Creation and Wrapping

The most important function is `sterrors.E(args ...interface{})`. It is used to create a new error.

As arguments, you can pass in any order the following types:

- error
- string (for passing the error message)

Apart from that, when a new error is created, the call position is remembered so that this information can
be used later in the call stack.
If you pass an error to this function, this error will be wrapped by the new error.

## Examples:

### Wrapping errors

You can simply wrap errors by passing the previous error to sterrors.E() as parameter.

In the stacktrace every wrapped error is an entry in the trace with its error message.
However, if you call Error() on the last error in the trace, it will return all error message in the form

{message}: {previous message1}: {previous message1}: ...

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/Bibob7/sterrors"
)

func main() {
	err := anotherMethod()

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

func anotherMethod() error {
	return sterrors.E("some error message")
}

```

Output:

```json
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
```

```
Simply print the wrapped error message: action not possible: some error message
Is initial error: true
```

Inspired by a talk from GopherCon 2019: https://www.youtube.com/watch?v=4WIhhzTTd0Y