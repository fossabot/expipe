// Copyright 2016 Arsham Shirvani <arshamshirvani@gmail.com>. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license
// License that can be found in the LICENSE file.

package internal

import (
	"io/ioutil"
	"strings"

	"github.com/Sirupsen/logrus"
)

type FieldLogger logrus.FieldLogger
type Level logrus.Level

type Logger struct{ *logrus.Logger }
type Entry struct{ *logrus.Entry }

func StandardLogger() *Logger { return &Logger{logrus.StandardLogger()} }

const (
	InfoLevel  = logrus.InfoLevel
	WarnLevel  = logrus.WarnLevel
	DebugLevel = logrus.DebugLevel
	ErrorLevel = logrus.ErrorLevel
)

// GetLogger returns the default logger with the given log level.
func GetLogger(level string) *Logger {
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

	return StandardLogger()
}

// DiscardLogger returns a dummy logger.
// This is useful for tests when you don't want to actually write to the Stdout.
func DiscardLogger() *Logger {
	log := logrus.New()
	log.Out = ioutil.Discard
	return &Logger{log}
}