package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

// Info ..
func Info(msg string) {
	log.Info().Msg(msg)
}

// Fatal ...
func Fatal(msg string, err error) {
	log.Fatal().Str("error", err.Error()).Msg(msg)
}

// Panic ...
func Panic(msg string, err error) {
	log.Panic().Msg(msg)
}
