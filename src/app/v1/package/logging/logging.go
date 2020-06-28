package logging

import (
	"os"

	"github.com/orn-id/orn-mn-boilerplate-go/src/utils/log"
	"github.com/orn-id/orn-mn-boilerplate-go/src/utils/log/local"
	"github.com/orn-id/orn-mn-boilerplate-go/src/utils/log/sentry"
)

// Logging ...
type Logging struct {
	Local  local.LogLocalInterface
	Sentry sentry.SentryInterface
}

// LoggingHandler ...
func LoggingHandler() *Logging {
	return &Logging{
		Local:  log.LogHandler().Local(),
		Sentry: log.LogHandler().Sentry(),
	}
}

// LoggingInterface ...
type LoggingInterface interface {
	Error(err error)
}

func (log *Logging) Error(err error) {
	if os.Getenv("APP_ENV") == "development" {
		log.Local.Exception(err)
	}
	log.Sentry.CaptureException(err)
}
