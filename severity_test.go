package sterrors

import (
	"encoding/json"
	"fmt"
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

			highestSeverity := HighestSeverity(lastErr)
			if tc.ExpectedSeverity != highestSeverity {
				t.Errorf("exptected severity %s is not actual severity %s", tc.ExpectedSeverity, highestSeverity)
			}
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
			if tc.ExpectedString != tc.Severity.String() {
				t.Errorf("exptected severity string %s is not actual severity string %s", tc.ExpectedString, tc.Severity.String())
			}
		})
	}
}

func TestSeverity_MarshalJSON(t *testing.T) {
	severity := SeverityError
	severityJson, err := json.Marshal(severity)

	if err != nil {
		t.Errorf("unable to marshal: %v", err)
	}

	expectedJson := "\"error\""
	actualJson := string(severityJson)
	if actualJson != expectedJson {
		t.Errorf("expected json \"%s\" is not actual json \"%s\"", expectedJson, actualJson)
	}
}
