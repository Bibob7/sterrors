package sterrors

import (
	"runtime"
	"strings"
)

type Caller struct {
	FuncName string `json:"funcName,omitempty"`
	File     string `json:"file,omitempty"`
	Line     int    `json:"line,omitempty"`
}

type CallStackEntry struct {
	ErrMessage string   `json:"msg,omitempty"`
	Severity   Severity `json:"severity,omitempty"`
	Caller     *Caller  `json:"caller,omitempty"`
}

// CallStack returns the callstack
// It adds CallStack entries recursively based on error causes
// as long they also implement the Error interface
func CallStack(err error) []CallStackEntry {
	var stack []CallStackEntry
	if err == nil {
		return stack
	}

	e, ok := err.(Error)
	if !ok {
		return append(stack, CallStackEntry{ErrMessage: err.Error()})
	}

	res := []CallStackEntry{{ErrMessage: e.Message(), Caller: e.Caller(), Severity: e.Severity()}}

	if e.Cause() == nil {
		return res
	}

	res = append(res, CallStack(e.Cause())...)

	return res
}

func caller() *Caller {
	pc, file, line, _ := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	nameSegments := strings.Split(details.Name(), "/")
	funcName := details.Name()
	if len(nameSegments) > 0 {
		funcName = nameSegments[len(nameSegments)-1]
	}
	return &Caller{File: file, Line: line, FuncName: funcName}
}
