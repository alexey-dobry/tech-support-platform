package server

import (
	"net/http"

	"github.com/alexey-dobry/tech-support-platform/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleGetClientData() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientID := c.Param("client_id")

		var activeSession models.Manager

		data, err := s.database.Query(nil, "SELECT * FROM sessions WHERE client_id=$1", ClientID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
			s.logger.Errorf("Could not query data from database: %s", err)
			return
		}

		data.Next()
		err = data.Scan(&activeSession.ManagerID, &activeSession.IsFree, &activeSession.ClientID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}

		c.JSON(http.StatusOK, activeSession.ManagerID)
	}
}

func (s *Server) handleGetManagerData() gin.HandlerFunc {
	return func(c *gin.Context) {
		managerID := c.Param("manager_id")

		var activeSession models.Manager

		data, err := s.database.Query(nil, "SELECT * FROM sessions WHERE manager_id=$1", managerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
			s.logger.Errorf("Could not query data from database: %s", err)
			return
		}

		data.Next()
		err = data.Scan(&activeSession.ManagerID, &activeSession.IsFree, &activeSession.ClientID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}

		c.JSON(http.StatusOK, activeSession.ClientID)
	}
}

func (s *Server) handleAddNewManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		var manager models.EndRequest
		if err := c.BindJSON(&manager); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		_, err := s.database.Exec(nil, "INSERT INTO sessions (manager_id, is_free, client_id) VALUES ($1, $2, $3)", manager.ManagerID, true, 0)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
			s.logger.Errorf("Could not query data from database: %s", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	}
}

func (s *Server) handleAssignManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		DataFromBot := models.Request{}
		if err := c.BindJSON(&DataFromBot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		DataFromDB := models.Manager{}

		data := s.database.QueryRow(nil, "SELECT * FROM sessions WHERE is_free=1")

		err := data.Scan(&DataFromDB.ManagerID, &DataFromDB.IsFree, &DataFromDB.ClientID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
			s.logger.Errorf("Error scanning data: %s", err)
			return
		}

		_, err = s.database.Exec(nil, "UPDATE sessions SET is_free=0,client_id=$1 WHERE manager_id=$2", DataFromBot.ClientID, DataFromDB.ManagerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
			s.logger.Errorf("Error executing command: %s", err)
			return
		}

		c.JSON(http.StatusOK, DataFromDB.ManagerID)
	}
}

func (s *Server) handleFreeManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		DataFromBot := models.EndRequest{}
		if err := c.BindJSON(&DataFromBot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		_, err := s.database.Exec(nil, "UPDATE sessions SET is_free = TRUE, client_id=0 WHERE manager_id=$1", DataFromBot.ManagerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
			s.logger.Errorf("Error executing command: %s", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	}
}

func (s *Server) handleEndSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		managerID := c.Param("manager_id")

		_, err := s.database.Exec(nil, "DELETE FROM sessions WHERE manager_id=$1", managerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
			s.logger.Errorf("Error executing command: %s", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
