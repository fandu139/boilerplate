package sentry

import (
	"os"

	"github.com/getsentry/sentry-go"
)

// Sentry ...
type handler struct{}

// New ...
func New() Contract {
	return &handler{}
}

// Init ...
func (s *handler) init() error {
	client := sentry.ClientOptions{
		Dsn:   os.Getenv("SENTRY_LOG_DSN"),
		Debug: true,
	}
	return sentry.Init(client)
}

// CaptureException ...
func (s *handler) CaptureException(err error) {
	s.init()
	go sentry.CaptureException(err)
}

// CaptureMessage ...
func (s *handler) CaptureMessage(msg string) {
	s.init()
	go sentry.CaptureMessage(msg)
}

// CaptureEvent ...
func (s *handler) CaptureEvent(event *sentry.Event) {
	s.init()
	go sentry.CaptureEvent(event)
}
