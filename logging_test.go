package sterrors

import (
	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, formatter, defaultFormatter)
}

func TestLog(t *testing.T) {
	formatter := &TestLogFormatter{}
	SetLogger(formatter)

	err := E("initial err")

	Log(err)

	assert.Equal(t, formatter.ExpectedErr, err)
}
