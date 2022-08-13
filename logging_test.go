package sterrors

import (
	"reflect"
	"testing"
)

type TestCustomLogger struct {
	ExpectedErr error
}

func (t *TestCustomLogger) Log(err error) {
	t.ExpectedErr = err
}

func TestSetFormatter(t *testing.T) {
	customLogger := &TestCustomLogger{}
	SetLogger(customLogger)

	if !reflect.DeepEqual(customLogger, logger) {
		t.Errorf("customLogger is not equal to the set logger")
	}
}

func TestLog(t *testing.T) {
	formatter := &TestCustomLogger{}
	SetLogger(formatter)

	err := E("initial err")

	Log(err)

	if formatter.ExpectedErr != err {
		t.Errorf("expected err: %s is not actual err: %s", formatter.ExpectedErr, err)
	}
}
