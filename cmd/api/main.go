package main

import (
	"log"
	"os"
	"rest/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := server.NewServer()

	log.Printf("Server starting on port %s", port)
	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
