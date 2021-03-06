package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fandu139/boilerplate/src/config"
	"github.com/fandu139/boilerplate/src/router"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	startApp(*environment)
}

func startApp(env string) {
	engine := config.SetupEngine(env)
	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	router.LoadRouter(engine)
	engine.Run(serverString)
}
