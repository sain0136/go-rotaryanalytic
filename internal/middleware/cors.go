package middleware

import (
	"net/http"
	"os"
)

// CorsMiddleware handles CORS headers for the API
func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin (from env)
		origin := os.Getenv("CORS_ORIGIN")
		if origin == "" {
			origin = "http://localhost:5515"
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)

		// Allow specific headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Allow specific methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Allow credentials
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next(w, r)
	}
}

// AuthMiddleware checks for valid session cookie
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Skip auth check for login and logout routes
		if r.URL.Path == "/api/login" || r.URL.Path == "/api/logout" {
			next(w, r)
			return
		}

		cookie, err := r.Cookie("session")
		if err != nil || cookie.Value != "authenticated" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status": "error", "message": "Unauthorized"}`))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
