package server

import (
	"github.com/alexey-dobry/tech-support-platform/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Server struct {
	serverAddress string
	router        *gin.Engine
	logger        logger.Logger
	database      *pgx.Conn
}

func New(db *pgx.Conn, logger logger.Logger, port string) *Server {
	s := Server{
		serverAddress: port,
		router:        gin.Default(),
		logger:        logger,
		database:      db,
	}

	s.initRoutes()

	s.logger.Info("Server instance created")
	return &s
}

func (s *Server) Run() {
	s.logger.Fatal(s.router.Run(":" + s.serverAddress))
}
