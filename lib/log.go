package lib

import (
	"io/ioutil"
	"strings"

	"github.com/Sirupsen/logrus"
)

// GetLogger returns the default logger with the given log level.
func GetLogger(level string) *logrus.Logger {
	logrus.SetLevel(logrus.ErrorLevel)
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true
	switch strings.ToLower(level) {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.ErrorLevel)
	}

	return logrus.StandardLogger()
}

// DiscardLogger returns a dummy logger.
// This is useful for tests when you don't want to actually write to the Stdout.
func DiscardLogger() *logrus.Logger {
	log := logrus.New()
	log.Out = ioutil.Discard
	return log
}
