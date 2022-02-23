package sterrors

import (
	"github.com/sirupsen/logrus"
)

type LogrusFormatter struct{}

func (f *LogrusFormatter) Log(err error) {
	e, ok := err.(Error)
	if !ok {
		logrus.Error(err)
		return
	}
	entry := logrus.WithFields(logrus.Fields{"stack": CallStack(e)})
	switch HighestSeverity(e) {
	case SeverityWarning:
		entry.Warnf("%s: %v", e.Caller().FuncName, e)
	case SeverityInfo:
		entry.Infof("%s: %v", e.Caller().FuncName, e)
	case SeverityDebug:
		entry.Debugf("%s: %v", e.Caller().FuncName, e)
	default:
		entry.Errorf("%s: %v", e.Caller().FuncName, e)
	}
}
