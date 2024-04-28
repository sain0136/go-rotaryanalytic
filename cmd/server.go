package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sain0136/go-rotaryanalytic/handlers"
)

func main() {
	mode := getMode()
	// dont call these functions rather register them so ResponseWriter  and Request can be passed as arguments to them automatically
	http.Handle("/", handlers.HomeHandler(mode))
	http.Handle("/login", handlers.LoginPageHandler(mode))
	// TOD have gpt explain why this happenng and needs to be wrapped like this
	http.Handle("/logs", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1 // Default page
		}
		handlers.LogsTable(w, r, page).ServeHTTP(w, r)
	}))
	http.Handle("/getLogsPath", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetLogsPath(w, r).ServeHTTP(w, r)
	}))

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
