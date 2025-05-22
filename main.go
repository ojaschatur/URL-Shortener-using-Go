// package main

// import (
// 	"crypto/md5"
// 	"encoding/hex"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// type URL struct {
// 	ID string `json:"id"`
// 	OriginalURL string `json:"original_url"`
// 	ShortURL string `json:"short_url"`
// 	CreationDate time.Time `json:"creation_date"`
// }

// var urlDB = make(map[string]URL)

// func generateShortURL(OriginalURL string) string {
// 	hasher := md5.New()
// 	hasher.Write([]byte(OriginalURL)) //converts the original URL from string to bytes
// 	data := hasher.Sum(nil)
// 	hash := hex.EncodeToString(data)

// 	fmt.Println("Hash: ", hash)
// 	fmt.Println("Final Hash: ", hash[:8])
// 	return hash[:8]
// }

// func createURL(originalURL string) string {
// 	shortURL := generateShortURL(originalURL)
// 	id := shortURL
// 	urlDB[id] = URL{
// 		ID: id,
// 		OriginalURL: originalURL,
// 		ShortURL: shortURL,
// 		CreationDate: time.Now(),
// 	}
// 	return shortURL
// }

// func getURL(id string) (URL, error) {
// 	url, ok := urlDB[id]
// 	if !ok {
// 		return URL{}, errors.New("URL not found")
// 	}
// 	return url, nil
// }

// func RootPageURL(w http.ResponseWriter, r *http.Request)  {
// 	fmt.Fprintf(w, "Hello World")

// }

// func ShortURLHandler(w http.ResponseWriter, r *http.Request)  {
// 	var data struct {
// 		URL string `json:"url"`
// 	}
// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	shortURL_ := createURL(data.URL)
// 	// fmt.Fprintf(w, shortURL)
// 	response := struct {
// 		ShortURL string `json:"short_url"`
// 	} {ShortURL:shortURL_}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func RedirectURLHandler(w http.ResponseWriter, r *http.Request)  {
// 	id := r.URL.Path[len("/redirect/"):]
// 	url, err := getURL(id)
// 	if err != nil {
// 		http.Error(w, "invalid request", http.StatusNotFound)
// 	}
// 	http.Redirect(w, r, url.OriginalURL, http.StatusFound)

// }

// func main() {
// 	//fmt.Println("Starting URL Shortener...")
// 	//OriginalURL := "https://github.com/ojaschatur"
// 	//generateShortURL(OriginalURL)

// 	// Register the handler function to handle all requests to the root url ("/")
// 	http.HandleFunc("/", RootPageURL)
// 	http.HandleFunc("/shorten", ShortURLHandler)

// 	// Start the HTTP server on port 3000
// 	fmt.Println("Starting on port 3000...")
// 	err := http.ListenAndServe(":3000", nil)
// 	if (err != nil) {
// 		fmt.Println("Error on starting the server.", err)
// 	}
// }

//

package main

import (
	"fmt"
	"net/http"
	"url-shortener/db"
	"url-shortener/handlers"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.ConnectDB()

	http.HandleFunc("/shorten", handlers.ShortenHandler)

	fmt.Println("Server started at :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
