package utils

import (
	"time"
	"url-shortener/db"
)

const RateLimit = 5

func CheckRateLimit(ip string) (int, time.Time, error) {
	var count int
	var reset time.Time

	err := db.DB.QueryRow("SELECT request_count, reset_time FROM rate_limits WHERE ip = $1", ip).Scan(&count, &reset)
	if err != nil {
		// Insert new IP
		reset = time.Now().Add(1 * time.Minute)
		_, err := db.DB.Exec("INSERT INTO rate_limits (ip, request_count, reset_time) VALUES ($1, $2, $3)", ip, 1, reset)
		return 1, reset, err
	}

	if time.Now().After(reset) {
		reset = time.Now().Add(1 * time.Minute)
		_, _ = db.DB.Exec("UPDATE rate_limits SET request_count = 1, reset_time = $1 WHERE ip = $2", reset, ip)
		return 1, reset, nil
	}

	if count >= RateLimit {
		return count, reset, nil
	}

	count++
	_, _ = db.DB.Exec("UPDATE rate_limits SET request_count = $1 WHERE ip = $2", count, ip)
	return count, reset, nil
}
