# Sterrors

[![codecov](https://codecov.io/gh/Bibob7/sterrors/branch/main/graph/badge.svg?token=2LURD0VD9X)](https://codecov.io/gh/Bibob7/sterrors)
[![CircleCI](https://circleci.com/gh/Bibob7/sterrors/tree/main.svg?style=svg)](https://circleci.com/gh/Bibob7/sterrors/tree/main)

sterrors is a simple library which will provide error types with stack traces, severity levels and error kinds.

You can extend the provides BaseError type and add you own application specific attributes that you may want to log.

## Error Creation and Wrapping

The most important function is `sterrors.E(args ...interface{})`. It is used to create a new error.

As arguments you can pass in any order the following types:

- error
- sterrors.Severity
- sterrors.Kind
- string (for passing the error message)

Apart from that, when a new error is created, the call position is remembered so that this information can
be used later in the call stack.

Here is an example:

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
	seconds := sterrors.E("action not possible", sterrors.SeverityError, err)

	e, _ := seconds.(sterrors.Error)
	jsonStackTrace, _ := json.Marshal(e.CallStack())
	fmt.Printf("%s", jsonStackTrace)
}

func anotherMethod() error {
	return sterrors.E("some error message", sterrors.SeverityWarning)
}

```

Output:

```json
[
   {
      "ErrMessage":"action not possible",
      "Caller":{
         "FuncName":"main.main",
         "File":"/Users/test/Repositories/sterrors/main/main.go",
         "Line":13
      }
   },
   {
      "ErrMessage":"some error message",
      "Caller":{
         "FuncName":"main.anotherMethod",
         "File":"/Users/test/Repositories/sterrors/main/main.go",
         "Line":21
      }
   }
]
```
