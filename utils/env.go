package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type EnvValue string

const (
	DEV_MODE       EnvValue = "DEV_MODE"
	BASE_LOG_PATH  EnvValue = "BASE_LOG_PATH"
	MYSQL_HOST     EnvValue = "MYSQL_HOST"
	MYSQL_PORT     EnvValue = "MYSQL_PORT"
	MYSQL_USER     EnvValue = "MYSQL_USER"
	MYSQL_PASSWORD EnvValue = "MYSQL_PASSWORD"
	MYSQL_DB_NAME  EnvValue = "MYSQL_DB_NAME"
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

func GetConfig(val EnvValue) (string, error) {
	env, ok := os.LookupEnv(string(val))
	if !ok {
		return "", fmt.Errorf("environment variable %s not found", val)
	}
	return env, nil
}
