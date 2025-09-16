package app

import (
	"log"

	"github.com/alexey-dobry/tech-support-platform/internal/pkg/logger"
	"github.com/alexey-dobry/tech-support-platform/internal/services/session_service/internal/server"
	"github.com/jackc/pgx/v5"
)

type App interface {
	Run()
}

type app struct {
	server *server.Server
}

func New(db *pgx.Conn, logger logger.Logger, cfg server.Config) App {
	a := app{
		server: server.New(db, logger, cfg.Port),
	}

	log.Print("App instance created")
	return &a
}

func (a *app) Run() {
	a.server.Run()
}
