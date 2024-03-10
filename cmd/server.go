package main

import (
	"fmt"
	"net/http"

	"github.com/sain0136/go-rotaryanalytic/handlers"
)

func main() {

	http.Handle("/", handlers.HomeHandler())

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
