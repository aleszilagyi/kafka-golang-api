package logger

import "github.com/sirupsen/logrus"

type StandardLogger struct {
	*logrus.Logger
}

func NewLogger() *StandardLogger {
	baseLogger := logrus.New()
	standardLogger := &StandardLogger{baseLogger}
	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}
