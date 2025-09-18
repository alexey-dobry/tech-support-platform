package server

func (s *Server) initRoutes() {
	s.router.POST("/auth", s.handleLogin())

	s.logger.Info("server routes was initialized")
}
