package sterrors

import (
	"github.com/sirupsen/logrus"
)

func LogrusLog(err error) {
	stErr, ok := err.(Error)
	if !ok {
		logrus.Error(err)
		return
	}

	entry := logrus.WithFields(logrus.Fields{
		"stack": stErr.CallStack(),
		"kind":  stErr.Kind()})

	logError(stErr, entry)
}

func logError(err Error, entry *logrus.Entry) {
	switch Level(err) {
	case SeverityWarning:
		entry.Data["severity"] = "warning"
		entry.Warnf("%s: %v", err.Caller().FuncName, err)
	case SeverityInfo:
		entry.Data["severity"] = "info"
		entry.Infof("%s: %v", err.Caller().FuncName, err)
	case SeverityDebug:
		entry.Data["severity"] = "debug"
		entry.Debugf("%s: %v", err.Caller().FuncName, err)
	case SeverityNotice:
		entry.Data["severity"] = "notice"
		entry.Debugf("%s: %v", err.Caller().FuncName, err)
	case SeverityError:
	default:
		entry.Data["severity"] = "error"
		entry.Errorf("%s: %v", err.Caller().FuncName, err)
	}
}
