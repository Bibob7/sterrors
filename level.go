package sterrors

const (
	SeverityNotice  Severity = 1
	SeverityDebug   Severity = 2
	SeverityInfo    Severity = 3
	SeverityWarning Severity = 4
	SeverityError   Severity = 5
)

func Level(err error) Severity {
	e, ok := err.(*BaseError)
	if !ok {
		return SeverityError
	}

	if e.Cause != nil {
		causeSeverity := Level(e.Cause)
		if causeSeverity > e.Severity {
			return causeSeverity
		}
	}
	return e.Severity
}
