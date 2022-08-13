package sterrors

import (
	"testing"
)

type TestLogFormatter struct {
	ExpectedErr error
}

func (t *TestLogFormatter) Log(err error) {
	t.ExpectedErr = err
}

func TestSetFormatter(t *testing.T) {
	formatter := &TestLogFormatter{}
	SetLogger(formatter)

	if formatter != defaultLogger {
		t.Errorf("formatter is not equal to defailt formatter")
	}
}

func TestLog(t *testing.T) {
	formatter := &TestLogFormatter{}
	SetLogger(formatter)

	err := E("initial err")

	Log(err)

	if formatter.ExpectedErr != err {
		t.Errorf("expected err: %s is not actual err: %s", formatter.ExpectedErr, err)
	}
}
