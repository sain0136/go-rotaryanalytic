package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/sain0136/go-rotaryanalytic/handlers"
)

func main() {
	mode := getMode()
	http.Handle("/", handlers.HomeHandler(mode))
	http.Handle("/login", handlers.LoginPageHandler(mode))

	// 	http.Handle("/hello", handlers.HelloHandler())
	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}

func getMode() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "PROD"
	}
	dir := filepath.Dir(filename)
	abs, err := filepath.Abs(filepath.Join(dir, "..", ".env"))
	if err != nil {
		log.Print("Failed to resolve absolute path to .env file")
		return "PROD"
	}
	if err := godotenv.Load(abs); err != nil {
		log.Print("No .env file found")
		return "PROD"
	}
	_, exists := os.LookupEnv("DEV_MODE")
	if exists {
		return "DEV"
	}
	return "PROD"
}
