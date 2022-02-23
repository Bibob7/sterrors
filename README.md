# Sterrors

[![codecov](https://codecov.io/gh/Bibob7/sterrors/branch/main/graph/badge.svg?token=2LURD0VD9X)](https://codecov.io/gh/Bibob7/sterrors)
[![CircleCI](https://circleci.com/gh/Bibob7/sterrors/tree/main.svg?style=svg)](https://circleci.com/gh/Bibob7/sterrors/tree/main)

sterrors is a simple library which will provide error types with stack traces and severity levels.

You can extend the provides BaseError type and add you own application specific attributes that you may want to log.

## Error Creation and Wrapping

The most important function is `sterrors.E(args ...interface{})`. It is used to create a new error.

As arguments you can pass in any order the following types:

- error
- sterrors.Severity
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
	second := sterrors.E("action not possible", sterrors.SeverityError, err)
	
	jsonStackTrace, _ := json.Marshal(sterrors.CallStack(second))
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

If you want to use logrus logger you can do it like that:

```go
package main

import (
	"github.com/Bibob7/sterrors"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	sterrors.SetLogger(&sterrors.LogrusLogger{}) // this is not necessary, because LogrusLogger is the default logger
	err := anotherMethod()

	// second error that results from the first one
	secondErr := sterrors.E("action not possible", sterrors.SeverityError, err)

	sterrors.Log(secondErr)
}

func anotherMethod() error {
	return sterrors.E("some error message", sterrors.SeverityWarning)
}
```

Output:

```json
{
  "level": "error",
  "msg": "main.main: action not possible",
  "stack": [
    {
      "msg": "action not possible",
      "severity": "error",
      "caller": {
        "funcName": "main.main",
        "file": "/Users/root/Repositories/sterrors/examples/main.go",
        "line": 14
      }
    },
    {
      "msg": "some error message",
      "severity": "warning",
      "caller": {
        "funcName": "main.anotherMethod",
        "file": "/Users/root/Repositories/sterrors/examples/main.go",
        "line": 20
      }
    }
  ],
  "time": "2022-02-23T22:18:11+01:00"
}
```

Inspired by a talk from GopherCon 2019: https://www.youtube.com/watch?v=4WIhhzTTd0Y