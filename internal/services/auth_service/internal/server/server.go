package server

import (
	"log"

	"github.com/alexey-dobry/tech-support-platform/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Config struct {
	Port string `yaml:"port" validate:"required" env:"PORT" env-default:"8080"`
}

type Server struct {
	router   *gin.Engine
	logger   logger.Logger
	database *pgx.Conn
}

func New(db *pgx.Conn, logger logger.Logger) *Server {
	s := Server{
		router:   gin.Default(),
		logger:   logger,
		database: db,
	}

	s.initRoutes()

	log.Println("Server instance created")
	return &s
}

func (s *Server) Run(cfg *Config) {
	log.Fatal(s.router.Run(":" + cfg.Port))
}
