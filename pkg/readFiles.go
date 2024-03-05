package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type LogEntry struct {
	UniqueId    string    `json:"uniqueId"`
	TypeOfLog   TypeOfLog `json:"typeOfLog"`
	Type        string    `json:"type"`
	LoginStatus string    `json:"loginStatus"`
	User        User      `json:"user"`
	Timestamp   string    `json:"timestamp"`
}

type TypeOfLog struct {
	Type        string `json:"type"`
	LoginStatus string `json:"loginStatus"`
}

type User struct {
	UserId int    `json:"userId"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}

func GetEnvStatus() (string, error) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Print("No .env file found")
		return "Error", fmt.Errorf("no .env file found")
	}
	path, exists := os.LookupEnv("BASE_LOG_PATH")

	if exists {
		return path, nil
	}
	return "No path found", nil
}

func ReadLogFile(path string) (string, error) {
	logs, err := os.ReadFile(path)
	if err != nil {
		errorMessage := fmt.Errorf("error return: %w", err)
		log.Print(errorMessage)
		return "", errorMessage
	}
	return string(logs), nil
}
