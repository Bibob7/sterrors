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
