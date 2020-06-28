package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/orn-id/orn-mn-boilerplate-go/src/middleware"
)

// SetupEngine server router configuration
func SetupEngine(env string) *gin.Engine {
	defaultMiddleware := middleware.DefaultMiddleware{}
	configEnvironment(env)
	router := gin.Default()
	router.Use(defaultMiddleware.CORSMiddleware())
	return router
}

// configEnvironment ...
func configEnvironment(env string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
	err = os.Setenv("APP_ENV", env)
	if err != nil {
		log.Println("Cannot Load APP_ENV")
	}
	switch env {
	case "development":
		log.Println("Development Engine Running")
	case "production":
		gin.SetMode(gin.ReleaseMode)
		log.Println("Engine Running")
	default:
		log.Println("Unknown Environment Engine Running")
	}

}
