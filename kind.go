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
	e, ok := err.(Error)
	if !ok {
		return kind == KindUnexpected
	}

	return e.Kind() == kind
}
