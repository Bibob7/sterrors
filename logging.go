package sterrors

var defaultFormatter Logger = &LogrusLogger{}

type Logger interface {
	Log(err error)
}

func SetLogger(formatter Logger) {
	defaultFormatter = formatter
}

func Log(err error) {
	defaultFormatter.Log(err)
}
