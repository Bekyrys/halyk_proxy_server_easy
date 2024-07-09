package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

var requestStore sync.Map

type ProxyRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type ProxyResponse struct {
	ID      string            `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int               `json:"length"`
}

func main() {
	http.HandleFunc("/proxy", handleProxy)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	var proxyReq ProxyRequest
	log.Println("Received request")

	if err := json.NewDecoder(r.Body).Decode(&proxyReq); err != nil {
		log.Printf("Failed to decode request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация входных данных
	if proxyReq.Method == "" || proxyReq.URL == "" {
		log.Println("Method and URL are required")
		http.Error(w, "Method and URL are required", http.StatusBadRequest)
		return
	}

	// Создание уникального ID для запроса
	requestID := uuid.New().String()
	log.Printf("Generated request ID: %s", requestID)

	// Создание нового запроса
	client := &http.Client{}
	req, err := http.NewRequest(proxyReq.Method, proxyReq.URL, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Установка заголовков
	for key, value := range proxyReq.Headers {
		req.Header.Set(key, value)
	}

	// Отправка запроса к стороннему сервису
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Чтение ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Сохранение запроса и ответа в памяти
	requestStore.Store(requestID, string(body))
	log.Printf("Stored request ID: %s", requestID)

	// Формирование ответа клиенту
	proxyResp := ProxyResponse{
		ID:     requestID,
		Status: resp.StatusCode,
		Headers: func(h http.Header) map[string]string {
			headers := make(map[string]string)
			for k, v := range h {
				headers[k] = v[0]
			}
			return headers
		}(resp.Header),
		Length: len(body),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(proxyResp); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
