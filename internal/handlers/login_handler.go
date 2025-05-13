package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/sain0136/go-rotaryanalytic/internal/config"
)

// TestHandler returns a simple JSON response for testing purposes
func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello from Go backend!"}`))
}

// LoginRequest represents the structure of the login request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the structure of the login response
type LoginResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// LoginHandler handles user authentication
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	adminUser := os.Getenv("ADMIN_USERNAME")
	if adminUser == "" {
		adminUser = "admin"
	}
	adminPass := os.Getenv("ADMIN_PASSWORD")
	if adminPass == "" {
		adminPass = "MyrotaryProjects_123"
	}

	if loginReq.Username == adminUser && loginReq.Password == adminPass {
		w.Header().Set("Content-Type", "application/json")

		expirationTime := time.Now().Add(config.SessionDuration)
		cookie := &http.Cookie{
			Name:     "session",
			Value:    "authenticated",
			Expires:  expirationTime,
			Path:     "/",
			HttpOnly: true,
			Secure:   config.CookieSecure,
			SameSite: getSameSite(config.CookieSameSite),
			Domain:   config.CookieDomain,
		}
		http.SetCookie(w, cookie)

		response := LoginResponse{
			Status:  "success",
			Message: "Login successful",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(LoginResponse{
		Status:  "error",
		Message: "Invalid credentials",
	})
}

// LogoutHandler handles user logout
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Delete the session cookie
	cookie := &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour), // Expire immediately
		HttpOnly: true,
		Secure:   config.CookieSecure,
		SameSite: getSameSite(config.CookieSameSite),
		Domain:   config.CookieDomain,
	}
	http.SetCookie(w, cookie)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{
		Status:  "success",
		Message: "Logged out successfully",
	})
}

func getSameSite(mode string) http.SameSite {
	switch mode {
	case "Strict":
		return http.SameSiteStrictMode
	case "None":
		return http.SameSiteNoneMode
	default:
		return http.SameSiteLaxMode
	}
}
