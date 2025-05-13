package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sain0136/go-rotaryanalytic/internal/config"
	"github.com/sain0136/go-rotaryanalytic/internal/routes"
)

func main() {
	// Load .env file in development (ignore error if not present)
	_ = godotenv.Load()
	config.LoadConfigFromEnv()

	// Setup routes
	routes.SetupRoutes()

	// Get port from env, default to 3000
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}

	// Start server
	fmt.Printf("Server starting on :%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
