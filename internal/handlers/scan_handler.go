package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/raj921/javascript-file-scanner/internal/scanner"
	"github.com/raj921/javascript-file-scanner/internal/utils"
)

type ScanResult struct {
	URL       string              `json:"url"`
	FoundURLs map[string][]string `json:"foundUrls"`
}

func ScanHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	content, err := utils.FetchContent(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch URL content: %v", err), http.StatusInternalServerError)
		return
	}

	foundURLs := scanner.ExtractURLs(content, url)

	result := ScanResult{
		URL:       url,
		FoundURLs: foundURLs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
