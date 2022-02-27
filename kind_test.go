package sterrors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIs(t *testing.T) {
	testCases := map[string]struct {
		IsUnknownErrorType bool
		IsKind             Kind
		CheckKind          Kind
		Result             bool
	}{
		"Is unknown error type": {
			IsUnknownErrorType: true,
			CheckKind:          KindUnexpected,
			Result:             true,
		},
		"Is unexpected": {
			IsKind:    KindUnexpected,
			CheckKind: KindUnexpected,
			Result:    true,
		},
		"Is not unexpected": {
			IsKind:    KindNotFound,
			CheckKind: KindUnexpected,
			Result:    false,
		},
		"Is not found": {
			IsKind:    KindNotFound,
			CheckKind: KindNotFound,
			Result:    true,
		},
		"Is not not found": {
			IsKind:    KindInvalidInput,
			CheckKind: KindNotFound,
			Result:    false,
		},
		"Is not allowed": {
			IsKind:    KindNotAllowed,
			CheckKind: KindNotAllowed,
			Result:    true,
		},
		"Is not not allowed": {
			IsKind:    KindAlreadyExists,
			CheckKind: KindNotAllowed,
			Result:    false,
		},
		"Is already exists": {
			IsKind:    KindAlreadyExists,
			CheckKind: KindAlreadyExists,
			Result:    true,
		},
		"Is not already exists": {
			IsKind:    KindUnexpected,
			CheckKind: KindAlreadyExists,
			Result:    false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			var err error
			if tc.IsUnknownErrorType {
				err = errors.New("unknown error")
			} else {
				err = E(tc.IsKind)
			}

			assert.Equal(t, tc.Result, Is(err, tc.CheckKind))
		})
	}
}
