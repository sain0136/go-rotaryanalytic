package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/sain0136/go-rotaryanalytic/handlers"
	"github.com/sain0136/go-rotaryanalytic/utils"
)

func main() {
	if err := utils.LoadEnv(); err != nil {
		panic(err)
	}
	// Registering the handlers
	http.Handle("/", handlers.HomeHandler())
	http.Handle("/login", handlers.LoginPageHandler())
	// TODO have gpt explain why this happenng and needs to be wrapped like this
	http.Handle("/logs", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1 // Default page to look for pagination logs
		}
		handlers.LogsTable(w, r, page).ServeHTTP(w, r)
	}))
	http.Handle("/getLogsPath", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetLogsPath(w, r).ServeHTTP(w, r)
	}))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
