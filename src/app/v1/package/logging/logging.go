package logging

import (
	"os"

	"github.com/sofyan48/boilerplate/src/utils/log"
	"github.com/sofyan48/boilerplate/src/utils/log/local"
	"github.com/sofyan48/boilerplate/src/utils/log/sentry"
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
