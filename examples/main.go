package main

import (
	"encoding/json"
	"fmt"
	"github.com/Bibob7/sterrors"
)

func main() {
	err := secondMethod()

	// second error that results from the first one
	second := sterrors.E("action not possible", sterrors.SeverityError, err)

	jsonStackTrace, _ := json.Marshal(sterrors.CallStack(second))
	fmt.Printf("%s", jsonStackTrace)
}

func secondMethod() error {
	return sterrors.E("some error message", sterrors.SeverityWarning)
}
