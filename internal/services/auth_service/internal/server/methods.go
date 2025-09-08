package server

import (
	"net/http"

	"github.com/alexey-dobry/tech-support-platform/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleGetLoginData() gin.HandlerFunc {
	return func(c *gin.Context) {
		DataFromBot := models.LoginData{}
		if err := c.BindJSON(&DataFromBot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			s.logger.Error("Could not bind json")
			return
		}

		data, err := s.database.Query(nil, "SELECT * FROM managers WHERE username=$1", DataFromBot.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error retreaving data from database"})
			s.logger.Errorf("Could not querry data from database: %s", err)
			return
		}

		DataFromDB := models.LoginData{}

		data.Next()
		err = data.Scan(&DataFromDB.Username, &DataFromDB.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding data from DB"})
			s.logger.Error(err)
			return
		}
		// Проверяем логин и пароль
		if DataFromBot.Password == DataFromDB.Password {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		}
	}
}
