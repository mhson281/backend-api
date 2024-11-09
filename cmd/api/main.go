package main

import (
	"net/http"
	"log"

	"github.com/mhson281/backend-api/internal/handlers"
)

func main() {
	// create a new http mux router
	mux := http.NewServeMux()

	mux.HandleFunc("/calculate", handlers.HandleCalculation)

	// start the HTTP router
	log.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}

}
