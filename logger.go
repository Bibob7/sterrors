package sterrors

import (
	"log"
)

type DefaultLogger struct{}

func (f *DefaultLogger) Log(err error) {
	log.Printf("%v", err)
}
