package sterrors

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

	res := []CallStackEntry{{ErrMessage: e.Error(), Caller: e.Caller()}}

	if e.Cause() == nil {
		return res
	}

	res = append(res, CallStack(e.Cause())...)

	return res
}
