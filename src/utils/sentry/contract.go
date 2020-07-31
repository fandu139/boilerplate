package sentry

import "github.com/getsentry/sentry-go"

// Contract ...
type Contract interface {
	CaptureException(err error)
	CaptureMessage(msg string)
	CaptureEvent(event *sentry.Event)
}
