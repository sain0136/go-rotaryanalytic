package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type envValue string

const (
	DEV_MODE      envValue = "DEV_MODE"
	BASE_LOG_PATH envValue = "BASE_LOG_PATH"
)

func LoadEnv() error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get current file path")
	}
	dir := filepath.Dir(filename)
	abs, err := filepath.Abs(filepath.Join(dir, "..", ".env"))
	if err != nil {
		log.Print("Failed to resolve absolute path to .env file")
		return fmt.Errorf("failed to resolve absolute path to .env file")
	}
	if err := godotenv.Load(abs); err != nil {
		log.Print("No .env file found")
		return fmt.Errorf("no .env file found")
	}
	return nil
}

func GetConfig(val envValue) (string, error) {
	env, ok := os.LookupEnv(string(val))
	if !ok {
		return "", fmt.Errorf("Environment variable %s not found", val)
	}
	return env, nil
}
