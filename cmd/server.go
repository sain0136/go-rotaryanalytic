package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/sain0136/go-rotaryanalytic/pkg"
	"github.com/sain0136/go-rotaryanalytic/views"
)

func main() {

	filePath, err := pkg.GetEnvStatus()
	if err != nil {
		fmt.Println(err)
	}
	rawData, writeErr := pkg.ReadLogFile(filePath)
	var component templ.Component
	if writeErr != nil {
		component = views.Hello("John", filePath, string(writeErr.Error()))
	} else {
		component = views.Hello("John", filePath, rawData)
	}
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
