package main

import (
	"data-enricher-dispatcher/client"
	"data-enricher-dispatcher/config"
	"data-enricher-dispatcher/logger"
	"data-enricher-dispatcher/service"
)

const dotEnv = ".env"

func main() {
	logger := logger.NewLogger()
	cfg, err := config.NewConfig(dotEnv)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Configuration loaded successfully:", cfg)

	apiClient := client.NewAPIClientV2(cfg)

	dispatcher := service.NewDispatcher(apiClient, logger, cfg)
	if err := dispatcher.Start(); err != nil {
		logger.Fatal("Failed to start dispatcher:", err)
	}
}
