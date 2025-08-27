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

type LoginPayload struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// GET HANDLER
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		resp := Response{
			Message: "Hello from auth services.",
			Status:  http.StatusOK,
		}

		json.NewEncoder(w).Encode(resp)
	}

	// POST HANDLER
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var accountLoggedIn LoginPayload
		if err := json.Unmarshal(body, &accountLoggedIn); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		fmt.Printf("Account logged in: %+v\n", accountLoggedIn)

		w.Header().Set("Content-Type", "application/json")
		resp := Response{
			Message: "Payment successfully",
			Status:  http.StatusOK,
		}

		json.NewEncoder(w).Encode(resp)
	}

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

	http.HandleFunc("/auth", paymentHandler)

	fmt.Printf("Server listening on port %s...\n", port)

	serve := http.ListenAndServe(":"+port, nil)
	if serve != nil {
		fmt.Println("Error starting server:", serve)
		return
	}

}
