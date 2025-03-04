package main

import (
	"fmt"
	"log"
	"money-transfer/internal/handler"
	"money-transfer/internal/repository"
	"money-transfer/internal/service"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Simple welcome message
	fmt.Fprintf(w, "Welcome to Money Transfer Service")
}

func methodHandler(allowedMethod string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != allowedMethod {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

func main() {
	// Create dependencies
	accountRepo := repository.NewAccountRepository()
	transferService := service.NewTransferService(accountRepo)

	// Create handlers
	accountHandler := handler.NewAccountHandler(accountRepo)
	transferHandler := handler.NewTransferHandler(transferService)

	mux := http.NewServeMux()
	mux.HandleFunc("/", methodHandler(http.MethodGet, homeHandler))
	mux.HandleFunc("/account/create", methodHandler(http.MethodPost, accountHandler.CreateAccount))
	mux.HandleFunc("/account/balance", methodHandler(http.MethodGet, accountHandler.GetAccountBalance))
	mux.HandleFunc("/transfer", methodHandler(http.MethodPost, transferHandler.Transfer))

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
