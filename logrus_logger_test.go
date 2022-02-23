package sterrors

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestWriter struct {
	Content []byte
}

func (w *TestWriter) Write(p []byte) (n int, err error) {
	w.Content = p
	return len(p), nil
}

type TestOutput struct {
	Msg       string               `json:"msg"`
	Level     string               `json:"level"`
	CallStack []TestCallStackEntry `json:"stack"`
}

type TestCallStackEntry struct {
	ErrMessage string `json:"msg"`
	Severity   string `json:"severity"`
	Caller     Caller `json:"caller"`
}

func TestLogrusFormatter_Log(t *testing.T) {
	testCases := map[string]struct {
		FinalLogLevel Severity
	}{
		"Severity Error": {
			FinalLogLevel: SeverityError,
		},
		"Severity Warning": {
			FinalLogLevel: SeverityWarning,
		},
		"Severity Notice": {
			FinalLogLevel: SeverityNotice,
		},
		"Severity Info": {
			FinalLogLevel: SeverityInfo,
		},
		"Severity Debug": {
			FinalLogLevel: SeverityDebug,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			outputWriter := &TestWriter{}
			logrus.SetOutput(outputWriter)
			logrus.SetFormatter(&logrus.JSONFormatter{})

			err := E("initial error", SeverityWarning)
			secondErr := E("second error", err, tc.FinalLogLevel)

			formatter := LogrusLogger{}
			formatter.Log(secondErr)

			var output TestOutput
			unmarshalErr := json.Unmarshal(outputWriter.Content, &output)

			e := err.(Error)
			secondE := secondErr.(Error)

			assert.Nil(t, unmarshalErr)
			assert.Equal(t, "sterrors.TestLogrusFormatter_Log.func1: second error", output.Msg)
			assert.Equal(t, HighestSeverity(secondE).String(), output.Level)
			assert.Len(t, output.CallStack, 2)
			assert.Equal(t, TestCallStackEntry{
				ErrMessage: "second error",
				Severity:   secondE.Severity().String(),
				Caller: Caller{
					FuncName: secondE.Caller().FuncName,
					File:     secondE.Caller().File,
					Line:     secondE.Caller().Line,
				},
			}, output.CallStack[0])
			assert.Equal(t, TestCallStackEntry{
				ErrMessage: "initial error",
				Severity:   e.Severity().String(),
				Caller: Caller{
					FuncName: e.Caller().FuncName,
					File:     e.Caller().File,
					Line:     e.Caller().Line,
				},
			}, output.CallStack[1])
		})
	}
}
