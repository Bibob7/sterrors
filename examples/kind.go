package main

import (
	"github.com/Bibob7/sterrors"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	sterrors.SetLogger(&sterrors.LogrusLogger{}) // this is not necessary, because LogrusLogger is the default logger

	// second error that results from the first one
	err := sterrors.E("action not possible", sterrors.KindNotAllowed)

	if sterrors.Is(err, sterrors.KindNotFound) {
		sterrors.Log(err)
	}
}
