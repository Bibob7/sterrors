package main

import (
	"encoding/json"
	"fmt"
	"github.com/Bibob7/sterrors"
)

func main() {
	err := anotherMethod()

	// second error that results from the first one
	secondErr := sterrors.E("action not possible", sterrors.SeverityError, err)

	jsonStackTrace, _ := json.Marshal(sterrors.CallStack(secondErr))
	fmt.Printf("%s", jsonStackTrace)
}

func anotherMethod() error {
	return sterrors.E("some error message", sterrors.SeverityWarning)
}
