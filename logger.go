package sterrors

import (
	"log"
)

type defaultLogger struct {
	logger *log.Logger
}

func NewDefaultLogger(logger *log.Logger) Logger {
	return &defaultLogger{
		logger: logger,
	}
}

func (f *defaultLogger) Log(err error) {
	log.Printf("%v", err)
}
