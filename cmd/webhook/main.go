package main

import (
	"log"
	"net/http"

	"github.com/livestorm/linear-workflows-manager/internal/webhook"
)

func main() {
	router, err := webhook.New()
	if err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}

	log.Print("Starting server on port :8080")
	err = http.ListenAndServe(":8080", router.Mux)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
