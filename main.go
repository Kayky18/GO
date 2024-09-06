package main

import (
	"github.com/kayky18/gopportunities/config"
	"github.com/kayky18/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	//Validate config
	if err != nil {
		logger.Errorf("Config err: %v", err)
		return
	}

	//iniando o router
	router.Initialize()
}
