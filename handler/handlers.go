package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/UHERO/go-url-shortener/shortener"
	"github.com/UHERO/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func CreateShortUrl(db *sql.DB, c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	var exists bool
	row := db.QueryRow("SELECT EXISTS(SELECT * FROM url_mapping WHERE original_url = ?)", creationRequest.LongUrl)
	if err := row.Scan(&exists); err != nil {
		fmt.Printf("CreateShortUrl Err. Failed to scan for existing row: %s", err)
	} else if !exists {
		shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl)
		store.SaveUrlMapping(db, shortUrl, creationRequest.LongUrl)

		host := "http://localhost:9808/"
		c.JSON(200, gin.H{
			"message":   "short url created successfully",
			"short_url": host + shortUrl,
		})

	}

}

func HandleShortUrlRedirect(db *sql.DB, c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl, err := store.RetrieveInitialUrl(db, shortUrl)
	if err != nil {
		fmt.Printf("Error retrieving initial Url")
	}
	c.Redirect(302, initialUrl)
}
