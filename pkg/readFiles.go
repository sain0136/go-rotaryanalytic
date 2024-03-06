package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type LogEntry struct {
	Level     int       `json:"level"`
	Time      int64     `json:"time"`
	Pid       int       `json:"pid"`
	Hostname  string    `json:"hostname"`
	RotaryLog RotaryLog `json:"rotaryLog"`
}

type RotaryLog struct {
	UniqueId  string `json:"uniqueId"`
	Type      string `json:"type"`
	Event     string `json:"event"`
	Status    string `json:"status"`
	Source    string `json:"source"`
	Target    string `json:"target"`
	Message   string `json:"message"`
	TimeStamp string `json:"timeStamp"`
}

// String returns a string representation of the LogEntry struct.
//
// We manually format the string representation instead of using the %+v verb
// in fmt.Sprintf because %+v would call the String method of the struct if it exists,
// leading to a recursive call in this case. By manually formatting the string,
// we avoid this recursion.
//
// The String method is useful when we want to print or log the contents of a LogEntry.
// It's also used when we convert a LogEntry to a string using the fmt package's printing functions.
func (l LogEntry) String() string {
	return fmt.Sprintf("Level: %d, Time: %d, Pid: %d, Hostname: %s, RotaryLog: %+v",
		l.Level, l.Time, l.Pid, l.Hostname, l.RotaryLog)
}

func (r RotaryLog) String() string {
	return fmt.Sprintf("UniqueId: %s, Type: %s, Event: %s, Status: %s, Source: %s, Target: %s, Message: %s, TimeStamp: %s",
		r.UniqueId, r.Type, r.Event, r.Status, r.Source, r.Target, r.Message, r.TimeStamp)
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

func ReadLogFile(path string) ([]RotaryLog, error) {
	logs, err := os.Open(path)
	if err != nil {
		errorMessage := fmt.Errorf("error return: %w", err)
		log.Print(errorMessage)
		return nil, errorMessage
	}
	defer logs.Close()

	scanner := bufio.NewScanner(logs)
	var rawLogs []string
	for scanner.Scan() {
		res := scanner.Text()
		rawLogs = append(rawLogs, res)
	}
	if err := scanner.Err(); err != nil {
		errorMessage := fmt.Errorf("error return: %w", err)
		return nil, errorMessage
	}
	return marshalLogs(rawLogs)
}

func marshalLogs(logs []string) ([]RotaryLog, error) {
	var entries []RotaryLog
	for _, logE := range logs {
		var entry LogEntry
		err := json.Unmarshal([]byte(logE), &entry)
		if err != nil {
			errorMessage := fmt.Errorf("error return: %w", err)
			log.Print(errorMessage)
			return nil, errorMessage
		}
		entries = append(entries, entry.RotaryLog)
	}
	return entries, nil
}
