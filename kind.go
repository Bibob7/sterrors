package sterrors

type Kind int

// Predefined kinds
const (
	KindUnexpected    Kind = 0
	KindInvalidInput  Kind = 1
	KindNotAllowed    Kind = 2
	KindAlreadyExists Kind = 3
	KindNotFound      Kind = 4
)

func Is(err error, kind Kind) bool {
	e, ok := err.(Error)
	if !ok {
		return kind == KindUnexpected
	}

	return e.Kind() == kind
}
