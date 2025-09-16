package server

import "log"

func (s *Server) initRoutes() {
	if s.router == nil {
		log.Fatal("Server router is not initialized")
	}

	s.router.POST("/create", s.handleAddNewManager())
	s.router.POST("/assign", s.handleAssignManager())
	s.router.POST("/end", s.handleFreeManager())
	s.router.GET("/sessions/manager/:manager_id", s.handleGetManagerData())
	s.router.GET("/sessions/client/:client_id", s.handleGetClientData())
	s.router.POST("/delete/:manager_id", s.handleEndSession())

	log.Print("server routes was initialized")
}
