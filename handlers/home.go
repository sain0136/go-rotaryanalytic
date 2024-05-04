// handlers/home.go
package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/sain0136/go-rotaryanalytic/pkg"
	"github.com/sain0136/go-rotaryanalytic/utils"
	"github.com/sain0136/go-rotaryanalytic/views"
)

var tokens = make(map[string]bool) // This map stores the valid tokens

func generateToken() string {
	b := make([]byte, 32)
	rand.Read(b) // Generates a random string for the cookie
	return base64.StdEncoding.EncodeToString(b)
}

type JsonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func HomeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var component templ.Component
		cookie, err := r.Cookie("token")
		if err != nil || !tokens[cookie.Value] {
			component = views.UnAthorized()
			templ.Handler(component).ServeHTTP(w, r)
			return
		}
		filePath, err := pkg.GetEnvStatus()
		if err != nil {
			fmt.Println(err)
		}
		rawLogEntries, writeErr, lastPage := pkg.ReadLogFile(filePath, 1)
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

func AuthorizeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create a new struct type to hold the JSON data
		type PasswordData struct {
			Password string `json:"password"`
		}

		// Decode the JSON from the request body
		var data PasswordData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Retrieve the password from the decoded JSON data
		password := data.Password
		// Retrieve the secret password from configuration
		secret, err := utils.GetConfig("PASSWORD_SECRET")
		// Set the Content-Type of the response
		w.Header().Set("Content-Type", "application/json")

		var response JsonResponse

		// Check for errors in retrieving the secret
		if err != nil {
			response = JsonResponse{
				Status:  "failure",
				Message: "Error retrieving configuration. Contact developer at jssr26@gmail.com.",
			}
		} else if secret != password {
			// Check if the provided password does not match the secret
			response = JsonResponse{
				Status:  "failure",
				Message: "Incorrect password.",
			}
		} else {
			// Password matches the secret
			token := generateToken()
			tokens[token] = true
			http.SetCookie(w, &http.Cookie{
				Name:     "token",
				Value:    token,
				HttpOnly: true,
			})
			response = JsonResponse{
				Status:  "success",
				Message: "Authorized successfully.",
			}
		}

		// Encode the response as JSON and handle errors
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Fatalf("Error encoding JSON: %v", err)
		}
	}
}
