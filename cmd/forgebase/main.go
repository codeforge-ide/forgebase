package main

import (
	"log"
	"os"

	"forgebase/internal/api"
	"forgebase/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	router := api.NewRouter(cfg)
	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	log.Printf("Forgebase server starting on :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
