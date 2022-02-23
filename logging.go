package sterrors

var defaultFormatter LogFormatter = &LogrusFormatter{}

type LogFormatter interface {
	Log(err error)
}

func SetFormatter(formatter LogFormatter) {
	defaultFormatter = formatter
}

func Log(err error) {
	defaultFormatter.Log(err)
}
