package main

import (
	"github.com/Bibob7/sterrors"
)

func main() {
	// second error that results from the first one
	err := sterrors.E("action not possible", sterrors.KindNotFound)

	if sterrors.Is(err, sterrors.KindNotFound) {
		sterrors.Log(err)
	}
}
