package main

import (
	"log"

	"github.com/alexey-dobry/tech-support-platform/internal/pkg/logger/zap"
	"github.com/alexey-dobry/tech-support-platform/internal/services/auth_service/internal/app"
	"github.com/alexey-dobry/tech-support-platform/internal/services/auth_service/internal/config"
	"github.com/alexey-dobry/tech-support-platform/internal/services/auth_service/internal/db"
)

func main() {
	cfg := config.MustLoad()

	logger := zap.NewLogger(cfg.Logger)

	db, err := db.New(cfg.DB)
	if err != nil {
		log.Fatal("Error creating database")
	}
	defer db.Close(nil)

	App := app.New(db, logger)

	App.Run(cfg)
}
