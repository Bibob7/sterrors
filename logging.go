package sterrors

import "log"

var logger = NewDefaultLogger(&log.Logger{})

type Logger interface {
	Log(err error)
}

func SetLogger(setLogger Logger) {
	logger = setLogger
}

func Log(err error) {
	logger.Log(err)
}
