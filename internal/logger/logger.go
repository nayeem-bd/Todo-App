package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger = logrus.New()

func init() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
}

// DefaultLogger returns the configured default logger
func DefaultLogger() *logrus.Logger {
	return logger
}

// SetLogLevel sets the log level for the logger
func SetLogLevel(level logrus.Level) {
	logger.SetLevel(level)
}

func Info(args ...interface{}) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Info(args...)
	}
}

func Error(args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Error(args...)
	}
}

func Debug(args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Debug(args...)
	}
}

func Warn(args ...interface{}) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Warn(args...)
	}
}

func Fatal(args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Fatal(args...)
		os.Exit(1) // Ensure the program exits after a fatal error
	}
}

func Panic(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	entry.Panic(args...)
}

func InfoWithFields(l interface{}, f logrus.Fields) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(f)
		entry.Info(l)
	}
}
