package routes

import (
	"net/http"

	"github.com/sain0136/go-rotaryanalytic/internal/handlers"
	"github.com/sain0136/go-rotaryanalytic/internal/middleware"
)

// SetupRoutes configures all routes for the application
func SetupRoutes() {
	// Apply both CORS and Auth middleware to all routes
	http.HandleFunc("/api/test", middleware.CorsMiddleware(middleware.AuthMiddleware(handlers.TestHandler)))
	http.HandleFunc("/api/login", middleware.CorsMiddleware(handlers.LoginHandler))
	http.HandleFunc("/api/logout", middleware.CorsMiddleware(handlers.LogoutHandler))
	http.HandleFunc("/api/getlogs", middleware.CorsMiddleware(middleware.AuthMiddleware(handlers.GetLogsHandler)))
}
