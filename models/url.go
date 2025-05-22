package models

import (
	"time"
	"url-shortener/db"
)

type URL struct {
	ID           int       `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
	ExpiryDate   time.Time `json:"expiry"`
}

func SaveURL(originalURL, shortURL string, expiry time.Time) error {
	_, err := db.DB.Exec("INSERT INTO urls (original_url, short_url, expiry_date) VALUES ($1, $2, $3)",
		originalURL, shortURL, expiry)
	return err
}

func GetURLByShortCode(code string) (URL, error) {
	row := db.DB.QueryRow("SELECT id, original_url, short_url, creation_date, expiry_date FROM urls WHERE short_url = $1", code)

	var u URL
	err := row.Scan(&u.ID, &u.OriginalURL, &u.ShortURL, &u.CreationDate, &u.ExpiryDate)
	return u, err
}
