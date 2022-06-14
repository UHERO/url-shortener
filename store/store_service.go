package store

import (
	"database/sql"
	"fmt"
)

type UrlMap struct {
	ID           int64
	original_url string
	short_url    string
}

var db *sql.DB

func SaveUrlMapping(db *sql.DB, shortUrl string, originalUrl string) sql.Result {
	result, err := db.Exec("INSERT INTO url_mapping (original_url, short_url) VALUES (?, ?)", originalUrl, shortUrl)
	if err != nil {
		panic(fmt.Sprintf(err.Error()))
	}
	return result
}

func RetrieveInitialUrl(db *sql.DB, shortUrl string) (string, error) {
	var urlMap UrlMap
	row := db.QueryRow("SELECT * FROM url_mapping WHERE short_url = ?", shortUrl)
	if err := row.Scan(&urlMap.ID, &urlMap.original_url, &urlMap.short_url); err != nil {
		if err == sql.ErrNoRows {
			return urlMap.original_url, fmt.Errorf("RetrieveInitialUrl %s: no url match", shortUrl)
		}
		return urlMap.original_url, fmt.Errorf("RetieveInitialUrl %s: %v", shortUrl, err)
	}
	return urlMap.original_url, nil
}
