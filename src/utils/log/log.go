package log

import (
	"github.com/orn-id/orn-mn-boilerplate-go/src/utils/log/local"
	"github.com/orn-id/orn-mn-boilerplate-go/src/utils/log/sentry"
)

// Log ...
type Log struct{}

// LogHandler ...
func LogHandler() *Log {
	return &Log{}
}

// LogInterface ...
type LogInterface interface {
	Local() *local.LogLocal
	Sentry() *sentry.Sentry
}

// Local ...
func (l *Log) Local() *local.LogLocal {
	return local.LogLocalHandler()
}

// Sentry ...
func (l *Log) Sentry() *sentry.Sentry {
	return sentry.SentryHandler()
}
