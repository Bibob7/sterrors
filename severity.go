package sterrors

import "encoding/json"

type Severity int

const (
	SeverityNotice Severity = iota
	SeverityDebug
	SeverityInfo
	SeverityWarning
	SeverityError
)

func (s Severity) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Severity) String() string {
	switch s {
	case SeverityNotice:
		return "notice"
	case SeverityDebug:
		return "debug"
	case SeverityInfo:
		return "info"
	case SeverityWarning:
		return "warning"
	default:
		return "error"
	}
}

// HighestSeverity returns the highest severity level for a whole stack of errors.
// If the error has no set severity level, the default is SeverityError
func HighestSeverity(err error) Severity {
	e, ok := err.(Error)
	if !ok {
		return SeverityError
	}

	if e.Cause() != nil {
		causeSeverity := HighestSeverity(e.Cause())
		if causeSeverity > e.Severity() {
			return causeSeverity
		}
	}
	return e.Severity()
}
