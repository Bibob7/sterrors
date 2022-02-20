package sterrors

import (
	"github.com/sirupsen/logrus"
)

func LogrusLog(err error) {
	sysErr, ok := err.(Error)
	if !ok {
		logrus.Error(err)
		return
	}

	entry := logrus.WithFields(logrus.Fields{
		"stack": sysErr.CallStack(),
		"kind":  sysErr.Kind()})

	logSysError(sysErr, entry)
}

func logSysError(sysErr Error, entry *logrus.Entry) {
	switch Level(sysErr) {
	case SeverityWarning:
		entry.Data["severity"] = "warning"
		entry.Warnf("%s: %v", sysErr.Caller().FuncName, sysErr)
	case SeverityInfo:
		entry.Data["severity"] = "info"
		entry.Infof("%s: %v", sysErr.Caller().FuncName, sysErr)
	case SeverityDebug:
		entry.Data["severity"] = "debug"
		entry.Debugf("%s: %v", sysErr.Caller().FuncName, sysErr)
	default:
		entry.Data["severity"] = "error"
		entry.Errorf("%s: %v", sysErr.Caller().FuncName, sysErr)
	}
}
