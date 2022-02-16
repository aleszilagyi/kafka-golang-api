package logger

import "github.com/sirupsen/logrus"

type Event struct {
	id      int
	message string
}

type StandardLogger struct {
	*logrus.Logger
}

func NewLogger() *StandardLogger {
	baseLogger := logrus.New()
	standardLogger := &StandardLogger{baseLogger}
	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

var (
	successMessage = Event{0, "Info: %s"}
	invalidArgMessage = Event{1, "Invalid arg: %s"}
	invalidArgValueMessage = Event{2, "Invalid value for arg: %s: %v"}
	missingArgMessage = Event{3, "Invalid arg: %s"}
	standardErrorMessage = Event{4, "Error: %s"}
)

func (l *StandardLogger) InvalidArg(argumentName string) {
  l.Errorf(invalidArgMessage.message, argumentName)
}

func (l *StandardLogger) InvalidArgValue(argumentName, argumentValue string) {
  l.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

func (l *StandardLogger) MissingArg(argumentName string) {
  l.Errorf(missingArgMessage.message, argumentName)
}