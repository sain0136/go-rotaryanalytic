package handlers

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/sain0136/go-rotaryanalytic/internal/config"
)

type LogEntry struct {
	Type          string          `json:"type"`
	Timestamp     string          `json:"timestamp"`
	Route         string          `json:"route"`
	RequestID     string          `json:"requestId"`
	Message       string          `json:"message"`
	CustomMessage string          `json:"customMessage"`
	Details       json.RawMessage `json:"details"`
	DetailsString string          `json:"detailsString"`
}

type GetLogsResponse struct {
	Logs       []LogEntry `json:"logs"`
	Page       int        `json:"page"`
	PageSize   int        `json:"pageSize"`
	Total      int        `json:"total"`
	TotalPages int        `json:"totalPages"`
}

// GetLogsHandler handles GET /api/getlogs with pagination
func GetLogsHandler(w http.ResponseWriter, r *http.Request) {
	logFilePath := config.LogFilePath
	file, err := os.Open(logFilePath)
	if err != nil {
		http.Error(w, "Log file not found or unreadable: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Pagination params
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if pageSize < 1 {
		pageSize = 50
	}

	// Read all lines
	scanner := bufio.NewScanner(file)
	var allLogs []LogEntry
	for scanner.Scan() {
		line := scanner.Bytes()
		var entry LogEntry
		if err := json.Unmarshal(line, &entry); err != nil {
			continue // skip malformed lines
		}
		// Marshal details to string for display
		if len(entry.Details) > 0 && string(entry.Details) != "null" {
			detailsStr, _ := json.Marshal(entry.Details)
			entry.DetailsString = string(detailsStr)
		}
		allLogs = append(allLogs, entry)
	}
	if err := scanner.Err(); err != nil {
		http.Error(w, "Error reading log file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	total := len(allLogs)
	totalPages := (total + pageSize - 1) / pageSize
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	logs := allLogs[start:end]

	resp := GetLogsResponse{
		Logs:       logs,
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
