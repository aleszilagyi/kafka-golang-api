package logger

import "github.com/sirupsen/logrus"

type StandardLogger struct {
	*logrus.Logger
}

func NewLogger() *StandardLogger {
	logger := logrus.New()
	standardLogger := &StandardLogger{logger}
	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}
