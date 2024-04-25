// handlers/home.go
package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/sain0136/go-rotaryanalytic/pkg"
	"github.com/sain0136/go-rotaryanalytic/views"
)

func HomeHandler(mode string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filePath, err := pkg.GetEnvStatus()
		if err != nil {
			fmt.Println(err)
		}
		rawLogEntries, writeErr := pkg.ReadLogFile(filePath)
		var component templ.Component
		if writeErr != nil {
			component = views.Hello("Peter Ivan Labelle", filePath, nil, mode)
		} else {
			component = views.Hello("Peter Ivan Labelle", filePath, rawLogEntries, mode)
		}
		templ.Handler(component).ServeHTTP(w, r)
	})
}

func LoginPageHandler(mode string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		component := views.LoginPage(mode)
		templ.Handler(component).ServeHTTP(w, r)
	})
}
