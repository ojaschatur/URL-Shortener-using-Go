package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"url-shortener/models"
	"url-shortener/utils"
)

func generateShortURL(originalURL string) string {
	return fmt.Sprintf("%x", time.Now().UnixNano())[:8] // simple unique hash
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	count, reset, err := utils.CheckRateLimit(ip)
	if err != nil || count > utils.RateLimit {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	var data struct {
		URL    string `json:"url"`
		Expiry int    `json:"expiry"` // minutes
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	short := generateShortURL(data.URL)
	expiry := time.Now().Add(time.Duration(data.Expiry) * time.Minute)

	err = models.SaveURL(data.URL, short, expiry)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"original_url":     data.URL,
		"short_url":        "http://localhost:3000/redirect/" + short,
		"expiry":           expiry,
		"rate_limit":       utils.RateLimit,
		"rate_limit_reset": reset,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
