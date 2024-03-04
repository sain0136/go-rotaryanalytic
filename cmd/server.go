package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/sain0136/go-rotaryanalytic/pkg"
	"github.com/sain0136/go-rotaryanalytic/views"
)

func main() {

	fileMessage, err := pkg.GetEnvStatus()
	if err != nil {
		fmt.Println(err)
	}
	component := views.Hello("John", fileMessage)
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
