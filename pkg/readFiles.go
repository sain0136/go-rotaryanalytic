package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/sain0136/go-rotaryanalytic/utils"
)

type LogEntry struct {
	Level     int       `json:"level"`
	Time      int64     `json:"time"`
	Pid       int       `json:"pid"`
	Hostname  string    `json:"hostname"`
	RotaryLog RotaryLog `json:"rotaryLog"`
}

type RotaryLog struct {
	UniqueId      string `json:"uniqueId"`
	Type          string `json:"type"`
	Event         string `json:"event"`
	Status        string `json:"status"`
	Source        string `json:"source"`
	Target        string `json:"target"`
	Message       string `json:"message"`
	CustomMessage string `json:"customMessage"`
	TimeStamp     string `json:"timeStamp"`
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
	return fmt.Sprintf("UniqueId: %s, Type: %s, Event: %s, Status: %s, Source: %s, Target: %s, Message: %s, CustomMessage: %s, TimeStamp: %s",
		r.UniqueId, r.Type, r.Event, r.Status, r.Source, r.Target, r.Message, r.CustomMessage, r.TimeStamp)
}

func GetEnvStatus() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Print("Failed to get current file path")
		return "Error", fmt.Errorf("failed to get current file path")
	}

	dir := filepath.Dir(filename) // get the directory of the current file
	log.Print("Current file path: ", dir)
	abs, err := filepath.Abs(filepath.Join(dir, "..", ".env")) // get the absolute path to the .env file
	if err != nil {
		log.Print("Failed to resolve absolute path to .env file")
		return "Error", fmt.Errorf("failed to resolve absolute path to .env file: %v", err)
	}
	log.Print("Found .env file at: ", abs)
	if err := godotenv.Load(abs); err != nil {
		log.Print("No .env file found")
		return "Error", fmt.Errorf("no .env file found")
	}
	path, exists := os.LookupEnv("BASE_LOG_PATH")

	if exists {
		return path, nil
	}
	return "No path found", nil
}

func ReadLogFile(path string, page int) ([]RotaryLog, int, error) {
	logs, err := os.Open(path)
	if err != nil {
		errorMessage := fmt.Errorf("error return: %w", err)
		log.Print(errorMessage)
		return nil, 0, errorMessage
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
		return nil, 0, errorMessage
	}

	marshaledLogs, err := marshalLogs(rawLogs)
	if err != nil {
		errorMessage := fmt.Errorf("error return: %w", err)
		return nil, 0, errorMessage
	}
	const logsPerPage = 10
	start := (page - 1) * logsPerPage
	end := start + logsPerPage
	lastPage := len(marshaledLogs) / logsPerPage
	// Useful GO syntax, its called slicing a slice
	return marshaledLogs[start:end], lastPage, nil
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
		time, err := utils.ParseToTimeWithUtc(entry.RotaryLog.TimeStamp)
		if err == nil {
			entry.RotaryLog.TimeStamp = time.UTC().Format("01/02/06 03:04:05 PM")
		}
		entries = append(entries, entry.RotaryLog)
	}
	reversed := utils.ReverseSlice(entries)
	return reversed, nil
}
