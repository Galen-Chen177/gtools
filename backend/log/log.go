package log

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type ILog interface {
	Log(string, string)
}

type Log struct{}

func NewLog() ILog {
	return &Log{}
}

func (l *Log) Log(level, in string) {
	switch strings.ToLower(level) {
	case "debug":
		logrus.Debug("vue ", in)
	case "info":
		logrus.Info("vue ", in)
	case "warn":
		logrus.Warn("vue ", in)
	case "error":
		logrus.Error("vue ", in)
	default:
		logrus.Warn("vue", in)
	}
}
