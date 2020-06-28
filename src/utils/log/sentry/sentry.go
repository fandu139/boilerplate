package sentry

import (
	"os"

	"github.com/getsentry/sentry-go"
)

// Sentry ...
type Sentry struct{}

// SentryHandler ...
func SentryHandler() *Sentry {
	return &Sentry{}
}

// SentryInterface ...
type SentryInterface interface {
	CaptureException(err error)
	CaptureMessage(msg string)
	CaptureEvent(event *sentry.Event)
}

// Init ...
func (logger *Sentry) init() error {
	client := sentry.ClientOptions{
		Dsn:   os.Getenv("SENTRY_LOG_DSN"),
		Debug: true,
	}
	return sentry.Init(client)
}

// CaptureException ...
func (logger *Sentry) CaptureException(err error) {
	logger.init()
	go sentry.CaptureException(err)
}

// CaptureMessage ...
func (logger *Sentry) CaptureMessage(msg string) {
	logger.init()
	go sentry.CaptureMessage(msg)
}

// CaptureEvent ...
func (logger *Sentry) CaptureEvent(event *sentry.Event) {
	logger.init()
	go sentry.CaptureEvent(event)
}
