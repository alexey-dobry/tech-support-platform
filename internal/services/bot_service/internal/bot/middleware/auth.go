package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Authenticate(username, password string) bool {

	payload := map[string]string{
		"username": username,
		"password": password,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println("Ошибка сериализации данных:", err)
		return false
	}

	resp, err := http.Post("http://localhost:8080/auth", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Ошибка запроса к микросервису:", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Ошибка аутентификации:", resp.Status)
		return false
	}

	// Парсим ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка чтения ответа:", err)
		return false
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("Ошибка парсинга ответа:", err)
		return false
	}

	return response["status"] == "success"
}
