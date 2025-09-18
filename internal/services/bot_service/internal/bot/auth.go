package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Authenticate(username, password, port string) bool {
	payload := map[string]string{
		"username": username,
		"password": password,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println("Ошибка сериализации данных:", err)
		return false
	}

	address := fmt.Sprintf("http://localhost:%s/auth", port)

	resp, err := http.Post(address, "application/json", bytes.NewBuffer(jsonData))
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
