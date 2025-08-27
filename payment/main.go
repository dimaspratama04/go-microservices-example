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

type RequestBody struct {
	Products []Product `json:"products"`
}

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Product struct {
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

// Handle success response
func httpSuccessHandler(w http.ResponseWriter, message string, statusCode int) {
	resp := Response{
		Message: message,
		Status:  statusCode,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

// Handle error response
func httpErrorHandler(w http.ResponseWriter, message string, statusCode int) {
	resp := Response{
		Message: message,
		Status:  statusCode,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)

	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		httpErrorHandler(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// GET HANDLER
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		resp := Response{
			Message: "Hello from payment services.",
			Status:  http.StatusOK,
		}

		json.NewEncoder(w).Encode(resp)
	}

	// POST HANDLER
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			httpErrorHandler(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var products []Product
		if err := json.Unmarshal(body, &products); err != nil {
			httpErrorHandler(w, "invalid JSON format: make sure use array payload", http.StatusBadRequest)
			return
		}

		fmt.Printf("Received products: %+v\n", products)

		httpSuccessHandler(w, "payment success.", http.StatusOK)
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

	http.HandleFunc("/", paymentHandler)

	fmt.Printf("Server listening on port %s...\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
