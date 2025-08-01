package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"forgebase/internal/api"
	"forgebase/internal/config"
)

func main() {
	// Load .env if present
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	router := api.NewRouter(cfg)
	log.Printf("Forgebase server starting on :%s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
