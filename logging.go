package sterrors

var defaultLogger Logger = &LogrusLogger{}

type Logger interface {
	Log(err error)
}

func SetLogger(logger Logger) {
	defaultLogger = logger
}

func Log(err error) {
	defaultLogger.Log(err)
}
