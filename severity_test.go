package sterrors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHighestLevel(t *testing.T) {
	testCases := map[string]struct {
		ExpectedSeverity         Severity
		ActualSeverityOccurrence []Severity
		InitialDefaultError      bool
	}{
		"Expected Severity Error": {
			ExpectedSeverity:         SeverityError,
			ActualSeverityOccurrence: []Severity{SeverityInfo, SeverityError, SeverityWarning},
		},
		"Expected Severity Info": {
			ExpectedSeverity:         SeverityInfo,
			ActualSeverityOccurrence: []Severity{SeverityInfo, SeverityDebug},
		},
		"Expected Severity Warning": {
			ExpectedSeverity:         SeverityWarning,
			ActualSeverityOccurrence: []Severity{SeverityWarning, SeverityInfo, SeverityDebug},
		},
		"Expected Severity Notice": {
			ExpectedSeverity:         SeverityNotice,
			ActualSeverityOccurrence: []Severity{SeverityNotice},
		},
		"Expected Severity Error because of unknown error type": {
			ExpectedSeverity:         SeverityError,
			ActualSeverityOccurrence: []Severity{SeverityNotice},
			InitialDefaultError:      true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			var lastErr error
			if tc.InitialDefaultError {
				lastErr = fmt.Errorf("some error")
			}
			for _, severity := range tc.ActualSeverityOccurrence {
				if lastErr != nil {
					lastErr = E(severity, lastErr)
				} else {
					lastErr = E(severity)
				}
			}

			assert.Equal(t, tc.ExpectedSeverity, HighestSeverity(lastErr))
		})
	}
}

func TestSeverity_String(t *testing.T) {
	testCases := map[string]struct {
		Severity       Severity
		ExpectedString string
	}{
		"Expected Severity Error": {
			Severity:       SeverityError,
			ExpectedString: "error",
		},
		"Expected Severity Info": {
			Severity:       SeverityInfo,
			ExpectedString: "info",
		},
		"Expected Severity Warning": {
			Severity:       SeverityWarning,
			ExpectedString: "warning",
		},
		"Expected Severity Notice": {
			Severity:       SeverityNotice,
			ExpectedString: "notice",
		},
		"Expected Severity Debug": {
			Severity:       SeverityDebug,
			ExpectedString: "debug",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.ExpectedString, tc.Severity.String())
		})
	}
}
