package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Структура для обработки JSON-данных
type EchoRequest struct {
	Message string `json:"message"`
}

type EchoResponse struct {
	Received string `json:"received"`
}

/*

func stringSum(message string, temp, sum int) string {
	for i := 0; i < len(message); i += 1 {
		temp := message[i]
		sum += int(temp)
	}
	message = message[sum:]
	return message
}

*/

func stringSum(message string) int {
	sum := 0
	for i := 0; i < len(message); i++ {
		// Преобразуем символ в цифру
		digit := message[i] - '0'
		// Проверяем, что это корректная цифра
		if digit < 0 || digit > 9 {
			continue // Игнорируем символы, не являющиеся цифрами
		}
		sum += int(digit)
	}
	return sum
}

func main() {
	http.HandleFunc("/api/echo", func(w http.ResponseWriter, r *http.Request) {
		// Проверка метода запроса
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		// Парсинг JSON-запроса
		var request EchoRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		/*
			stringSum(request.Message, temp, sum)
		*/

		sum := stringSum(request.Message)

		// Формирование JSON-ответа
		/*
			response := EchoResponse{
				Received: request.Message,
			}
		*/

		response := EchoResponse{
			Received: fmt.Sprintf("Sum of digits: %d", sum),
		}

		// Установка заголовка и отправка ответа
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	})

	// Запуск сервера
	port := 8080
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
