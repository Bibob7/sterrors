package sterrors

import (
	"log"
	"testing"
)

type TestWriter struct {
	Content []byte
}

func (w *TestWriter) Write(p []byte) (n int, err error) {
	w.Content = p
	return len(p), nil
}

func TestDefaultLogger(t *testing.T) {
	err := E("initial error", SeverityWarning)

	outputWriter := &TestWriter{}
	logger := NewDefaultLogger(log.New(outputWriter, "", 0))

	logger.Log(err)

	expectedOutput := ""
	actualOutput := string(outputWriter.Content)
	if actualOutput != expectedOutput {
		t.Errorf("expected logger output \"%s\" is not actual output \"%s\"", expectedOutput, actualOutput)
	}
}
