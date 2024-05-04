// handlers/home.go
package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/sain0136/go-rotaryanalytic/pkg"
	"github.com/sain0136/go-rotaryanalytic/utils"
	"github.com/sain0136/go-rotaryanalytic/views"
)

func HomeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filePath, err := pkg.GetEnvStatus()
		if err != nil {
			fmt.Println(err)
		}
		rawLogEntries, writeErr, lastPage := pkg.ReadLogFile(filePath, 1)
		var component templ.Component
		content := views.LogTable(rawLogEntries, lastPage)
		if writeErr != nil {
			component = views.Home(filePath, nil, lastPage, content)
		} else {
			component = views.Home(filePath, rawLogEntries, lastPage, content)
		}
		templ.Handler(component).ServeHTTP(w, r)
	})
}

func LoginPageHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mode, err := utils.GetConfig(utils.DEV_MODE)
		if err != nil {
			mode = "PROD"
		}
		if mode == "True" {
			mode = "DEV"
		} else {
			mode = "PROD"
		}
		component := views.LoginPage(mode)
		templ.Handler(component).ServeHTTP(w, r)
	})
}

func LogsTable(w http.ResponseWriter, r *http.Request, page int) http.Handler {
	filePath, err := pkg.GetEnvStatus()
	if err != nil {
		fmt.Println(err)
	}
	rawLogEntries, writeErr, lastPage := pkg.ReadLogFile(filePath, page)

	var component templ.Component
	if writeErr != nil {
		component = views.LogTable(nil, lastPage)
	} else {
		component = views.LogTable(rawLogEntries, lastPage)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(component).ServeHTTP(w, r)
	})
}

func GetLogsPath(w http.ResponseWriter, r *http.Request) http.Handler {
	filePath, err := pkg.GetEnvStatus()
	var path string
	if err != nil {
		path = "Logs was not found"
	} else {
		path = fmt.Sprintf("Logs found at: %s", filePath)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		component := views.LogPath(path)
		templ.Handler(component).ServeHTTP(w, r)
	})
}
