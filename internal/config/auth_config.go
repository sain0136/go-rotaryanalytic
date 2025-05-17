package config

import (
	"os"
	"time"
)

var (
	// SessionDuration is the duration for which the session cookie is valid
	SessionDuration time.Duration
	CookieSecure    bool
	CookieSameSite  string
	CookieDomain    string
	LogFilePath     string
)

// LoadConfigFromEnv loads the configuration from environment variables
func LoadConfigFromEnv() {
	// Session Duration
	dur := os.Getenv("SESSION_DURATION")
	if dur == "" {
		dur = "1m"
	}
	d, err := time.ParseDuration(dur)
	if err != nil {
		d = time.Minute
	}
	SessionDuration = d

	// Cookie Secure
	val := os.Getenv("COOKIE_SECURE")
	CookieSecure = val == "true" || val == "1"

	// SameSite
	val = os.Getenv("COOKIE_SAMESITE")
	if val == "" {
		val = "Lax"
	}
	CookieSameSite = val

	// Domain
	val = os.Getenv("COOKIE_DOMAIN")
	if val == "" {
		val = "localhost"
	}
	CookieDomain = val

	// Log File Path
	val = os.Getenv("LOG_FILE_PATH")
	if val == "" {
		val = "C://Users//jssr2//OneDrive//Documents//Github//myrotary-projects//backend//rotary.log"
	}
	LogFilePath = val
}
