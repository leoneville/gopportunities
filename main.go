package main

import (
	"os"

	"github.com/leoneville/gopportunities/config"
	"github.com/leoneville/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		os.Exit(1)
	}

	router.Initialize()
}
