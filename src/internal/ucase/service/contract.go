package service

import (
	"github.com/sofyan48/boilerplate/src/appctx/apps"
)

// Health ...
type Health interface {
	HealthStatus() (*apps.Result, *apps.Error)
	HealthCheck() (*apps.Result, *apps.Error)
}
