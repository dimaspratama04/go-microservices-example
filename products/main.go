package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Products struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity"`
}

// Set CORS headers
func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func httpResponseHandler(w http.ResponseWriter, message string, statusCode int) {
	resp := Response{
		Message: message,
		Status:  statusCode,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

func homeProductsHandler(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)

	if r.URL.Path != "/" {
		httpResponseHandler(w, "path not exist.", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodOptions && r.Method != http.MethodGet {
		httpResponseHandler(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		httpResponseHandler(w, "hello from products services.", http.StatusOK)
	}

}

func payProductsHandler(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)

	if r.Method != http.MethodPost {
		httpResponseHandler(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			httpResponseHandler(w, "failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var products []Products
		if err := json.Unmarshal(body, &products); err != nil {
			httpResponseHandler(w, "invalid JSON format.", http.StatusBadRequest)
			return
		}

		fmt.Printf("Received products: %+v\n", products)

		httpResponseHandler(w, "products payment sucessfully.", http.StatusOK)
	}
}

func main() {
	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	http.HandleFunc("/", homeProductsHandler)
	http.HandleFunc("/pay", payProductsHandler)

	fmt.Printf("Server listening on port %s...\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
