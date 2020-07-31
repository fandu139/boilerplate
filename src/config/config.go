package config

import (
	"io/ioutil"
	"os"

	"github.com/sofyan48/boilerplate/src/utils/log"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/sofyan48/boilerplate/src/middleware"
)

// SetupEngine server router configuration
func SetupEngine(env string) *gin.Engine {
	defaultMiddleware := middleware.DefaultMiddleware{}
	configEnvironment(env)
	router := gin.Default()
	router.Use(defaultMiddleware.CORSMiddleware())
	router.Use(defaultMiddleware.GetContext())

	ginlog := zerolog.New(os.Stdout).With().Logger()
	router.Use(logger.SetLogger(logger.Config{
		Logger: &ginlog,
		UTC:    true,
	}))
	return router
}

// configEnvironment ...
func configEnvironment(env string) {
	//setup log to JSON
	gin.DefaultWriter = ioutil.Discard

	// checking docekrize mode
	if env == "dockerize" {
		gin.DefaultWriter = ioutil.Discard
		log.Info("Docker Development Engine Running")
		return
	}

	// config not in dockerize mode
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", nil)
		os.Exit(1)
	}
	err = os.Setenv("APP_ENV", env)
	if err != nil {
		log.Info("Cannot Load APP_ENV")
	}

	// switching environment
	switch env {
	case "development":
		log.Info("Development Engine Running")

	case "production":
		// setup gin to release mode
		gin.SetMode(gin.ReleaseMode)
		log.Info("Engine Running")
	default:
		log.Info("Unknown Environment Engine Running")
	}
}
