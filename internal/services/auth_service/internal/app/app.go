package app

import (
	"log"

	"github.com/alexey-dobry/tech-support-platform/internal/pkg/logger"
	"github.com/alexey-dobry/tech-support-platform/internal/services/auth_service/internal/config"
	"github.com/alexey-dobry/tech-support-platform/internal/services/auth_service/internal/server"
	"github.com/jackc/pgx/v5"
)

type App struct {
	server *server.Server
}

func New(db *pgx.Conn, logger logger.Logger) *App {
	a := App{
		server: server.New(db, logger),
	}

	log.Print("App instance created")
	return &a
}

func (a *App) Run(cfg *config.Config) {
	log.Print("App is running...")
	a.server.Run(&cfg.Server)
}
