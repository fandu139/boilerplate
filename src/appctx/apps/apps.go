package apps

import "fmt"

// Error ...
type Error struct {
	Name     string
	Messages string
	Code     int
	Errors   error
	sentry   *SentryLog
}

// SentryLog ...
type SentryLog struct {
	DSN   string
	Debug bool
}

// Data ...
type Result struct {
	Name           string
	Messages       string
	Code           int
	Result         interface{}
	MessagesDetail map[string]string
	Meta           interface{}
}

// Sentry ...
func (e *Error) Sentry() *Error {
	sentrySetup := &SentryLog{}
	sentrySetup.DSN = ""
	sentrySetup.Debug = false
	e.sentry = sentrySetup
	return e
}

// Send ...
func (e *Error) Send() {
	fmt.Println("ERROR: ", e)
	fmt.Println("SENTRY: ", e.sentry)
}
