package sterrors

type Kind string

// Predefined kinds
const (
	KindUnexpected    Kind = "unexpected"
	KindInvalidInput  Kind = "invalid input"
	KindNotAllowed    Kind = "not allowed"
	KindAlreadyExists Kind = "already exists"
	KindNotFound      Kind = "not found"
)

func Is(err error, kind Kind) bool {
	if err == nil {
		return false
	}
	e, ok := err.(Error)
	if !ok {
		return kind == KindUnexpected
	}

	return e.Kind() == kind
}

func IsInStack(err error, kind Kind) bool {
	if Is(err, kind) {
		return true
	}

	e, ok := err.(Error)
	if !ok {
		return kind == KindUnexpected
	}

	if e.Cause() == nil {
		return false
	}

	return IsInStack(e.Cause(), kind)
}
