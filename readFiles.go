package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ReadLogFile() error {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return err
	}
	path, exists := os.LookupEnv("BASE_LOG_PATH")

	if exists {
		fmt.Println(path)
	}
	return nil

}
