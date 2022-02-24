package sterrors

import (
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct{}

func (f *LogrusLogger) Log(err error) {
	e, ok := err.(Error)
	if !ok {
		logrus.Error(err)
		return
	}
	entry := logrus.WithFields(logrus.Fields{"stack": CallStack(e)})
	switch HighestSeverity(e) {
	case SeverityInfo:
		entry.Infof("%s: %v", e.Caller().FuncName, e)
	case SeverityNotice:
		entry.Infof("%s: %v", e.Caller().FuncName, e)
	case SeverityWarning:
		entry.Warnf("%s: %v", e.Caller().FuncName, e)
	case SeverityDebug:
		entry.Debugf("%s: %v", e.Caller().FuncName, e)
	default:
		entry.Errorf("%s: %v", e.Caller().FuncName, e)
	}
}
