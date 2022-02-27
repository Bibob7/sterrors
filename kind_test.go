package sterrors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIs(t *testing.T) {
	testCases := map[string]struct {
		Error     error
		CheckKind Kind
		Result    bool
	}{
		"Is unknown error type": {
			Error:     errors.New("unknown error"),
			CheckKind: KindUnexpected,
			Result:    true,
		},
		"Error is nil": {
			Error:     nil,
			CheckKind: KindAlreadyExists,
			Result:    false,
		},
		"Is unexpected": {
			Error:     E(KindUnexpected),
			CheckKind: KindUnexpected,
			Result:    true,
		},
		"Is not unexpected": {
			Error:     E(KindNotFound),
			CheckKind: KindUnexpected,
			Result:    false,
		},
		"Is not found": {
			Error:     E(KindNotFound),
			CheckKind: KindNotFound,
			Result:    true,
		},
		"Is not not found": {
			Error:     E(KindInvalidInput),
			CheckKind: KindNotFound,
			Result:    false,
		},
		"Is not allowed": {
			Error:     E(KindNotAllowed),
			CheckKind: KindNotAllowed,
			Result:    true,
		},
		"Is not not allowed": {
			Error:     E(KindAlreadyExists),
			CheckKind: KindNotAllowed,
			Result:    false,
		},
		"Is already exists": {
			Error:     E(KindAlreadyExists),
			CheckKind: KindAlreadyExists,
			Result:    true,
		},
		"Is not already exists": {
			Error:     E(KindUnexpected),
			CheckKind: KindAlreadyExists,
			Result:    false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.Result, Is(tc.Error, tc.CheckKind))
		})
	}
}

func TestIsInStack(t *testing.T) {
	testCases := map[string]struct {
		Error     error
		CheckKind Kind
		Result    bool
	}{
		"Is unknown error type": {
			Error:     errors.New("unknown error"),
			CheckKind: KindUnexpected,
			Result:    true,
		},
		"Error is nil": {
			Error:     nil,
			CheckKind: KindAlreadyExists,
			Result:    false,
		},
		"Kind in stack": {
			Error:     E(KindNotAllowed, E(KindNotFound)),
			CheckKind: KindNotFound,
			Result:    true,
		},
		"Kind not in stack": {
			Error:     E(KindNotAllowed, E(KindNotFound)),
			CheckKind: KindInvalidInput,
			Result:    false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.Result, IsInStack(tc.Error, tc.CheckKind))
		})
	}
}
