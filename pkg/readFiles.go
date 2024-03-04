package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvStatus() (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return "Error", fmt.Errorf("no .env file found")
	}
	path, exists := os.LookupEnv("BASE_LOG_PATH")

	if exists {
		return path, nil
	}
	return "No path found", nil
}

func ReadLogFile() {

}
