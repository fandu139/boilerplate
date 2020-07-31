package contract

import "github.com/gin-gonic/gin"

type Health interface {
	HealthStatus(context *gin.Context)
	HealthCheck(context *gin.Context)
}
