package main

import (
	"github.com/Bibob7/sterrors"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	err := anotherMethod()

	// second error that results from the first one
	secondErr := sterrors.E("action not possible", sterrors.SeverityError, err)

	sterrors.Log(secondErr)
}

func anotherMethod() error {
	return sterrors.E("some error message", sterrors.SeverityWarning)
}
