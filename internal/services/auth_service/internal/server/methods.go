package server

import (
	"net/http"

	"github.com/alexey-dobry/tech-support-platform/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		DataFromBot := models.LoginData{}
		if err := c.BindJSON(&DataFromBot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			s.logger.Errorf("Could not bind json: %s", err)
			return
		}

		//Делаем запрос в бд, чтобы найти сущность по юзернейму
		data, err := s.database.Query(nil, "SELECT * FROM managers WHERE username=$1", DataFromBot.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
			s.logger.Errorf("Could not querry data from database: %s", err)
			return
		}

		DataFromDB := models.LoginData{}

		//Записываем данные из запроса в бд
		data.Next()
		err = data.Scan(&DataFromDB.Username, &DataFromDB.Password)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}

		// Проверяем логин и пароль
		if DataFromBot.Password == DataFromDB.Password {
			c.JSON(http.StatusOK, gin.H{})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{})
		}
	}
}
