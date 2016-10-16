package common

import (
	"log"
	"os"
)

// Logger is interface for output log
type Logger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Panicf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
}

var debugLogger = log.New(os.Stderr, "[DEBUG]", log.LstdFlags)
var infoLogger = log.New(os.Stderr, "[INFO]", log.LstdFlags)
var warnLogger = log.New(os.Stderr, "[WARN]", log.LstdFlags)
var errorLogger = log.New(os.Stderr, "[ERR]", log.LstdFlags)
var panicLogger = log.New(os.Stderr, "[PANIC]", log.LstdFlags)
var fatalLogger = log.New(os.Stderr, "[FATAL]", log.LstdFlags)

type defaultLogger struct{}

func (l defaultLogger) Debugf(format string, v ...interface{}) {
	debugLogger.Printf(format, v...)
}

func (l defaultLogger) Infof(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
}

func (l defaultLogger) Warnf(format string, v ...interface{}) {
	warnLogger.Printf(format, v...)
}

func (l defaultLogger) Errorf(format string, v ...interface{}) {
	errorLogger.Printf(format, v...)
}

func (l defaultLogger) Panicf(format string, v ...interface{}) {
	panicLogger.Panicf(format, v...)
}

func (l defaultLogger) Fatalf(format string, v ...interface{}) {
	fatalLogger.Fatalf(format, v...)
}

// DefaultLogger is logger with std out
var DefaultLogger = defaultLogger{}
