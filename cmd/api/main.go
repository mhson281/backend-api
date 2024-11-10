package main

import (
	"log"
	"net/http"

	"github.com/mhson281/backend-api/internal/database"
	"github.com/mhson281/backend-api/internal/handlers"
	"github.com/mhson281/backend-api/internal/middleware"
)

func main() {
	// initialize the sqlite db
	database.Init()

	// create a new http mux router
	mux := http.NewServeMux()
	mux.Handle("/calculate", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.HandleCalculation)))
	mux.HandleFunc("/register", handlers.HandleRegister)
  mux.HandleFunc("/login", handlers.HandleLogin)

	// start the HTTP router
	log.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}

}
