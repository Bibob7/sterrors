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
		entry.Data["severity"] = "warning"
		entry.Warnf("%s: %v", e.Caller().FuncName, e)
	case SeverityInfo:
		entry.Data["severity"] = "info"
		entry.Infof("%s: %v", e.Caller().FuncName, e)
	case SeverityDebug:
		entry.Data["severity"] = "debug"
		entry.Debugf("%s: %v", e.Caller().FuncName, e)
	default:
		entry.Data["severity"] = "error"
		entry.Errorf("%s: %v", e.Caller().FuncName, e)
	}
}
