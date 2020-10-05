package service

import (
	"github.com/fandu139/boilerplate/src/appctx/apps"
)

// Health ...
type Health interface {
	HealthStatus() (*apps.Result, *apps.Error)
	HealthCheck() (*apps.Result, *apps.Error)
}
