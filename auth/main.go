package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type LoginPayload struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	RequestId string `json:"request_id"`
	Message   string `json:"message"`
	Status    int    `json:"status"`
}

func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func httpResponseHandler(w http.ResponseWriter, message string, statusCode int) {
	resp := Response{
		RequestId: uuid.NewString(),
		Message:   message,
		Status:    statusCode,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

func homeAuthHandler(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)

	if r.URL.Path != "/" {
		httpResponseHandler(w, "path not exist.", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet && r.Method != http.MethodOptions {
		httpResponseHandler(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	httpResponseHandler(w, "hello from auth services.", http.StatusOK)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpResponseHandler(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		httpResponseHandler(w, "failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var accountLoggedIn LoginPayload
	if err := json.Unmarshal(body, &accountLoggedIn); err != nil {
		httpResponseHandler(w, "invalid JSON format.", http.StatusBadRequest)
		return
	}

	fmt.Printf("Account logged in: %+v\n", accountLoggedIn)

	w.Header().Set("Content-Type", "application/json")

	httpResponseHandler(w, "successfully logged in.", http.StatusOK)
}

func main() {
	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	http.HandleFunc("/", homeAuthHandler)
	http.HandleFunc("/login", loginHandler)

	fmt.Printf("Server listening on port %s...\n", port)

	serve := http.ListenAndServe(":"+port, nil)
	if serve != nil {
		fmt.Println("Error starting server:", serve)
		return
	}

}
