package store

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortURL := "ZV7FREWM"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockResult := sqlmock.NewRows([]string{"ID", "original_url", "short_url"}).AddRow('1', initialLink, shortURL)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM url_mapping WHERE short_url = ?")).WillReturnRows(mockResult)

	// Retrieve initial URL
	retrievedUrl, err := RetrieveInitialUrl(db, shortURL)
	if err != nil {
		panic(err)
	}
	fmt.Println(retrievedUrl)
	assert.Equal(t, initialLink, retrievedUrl)
}
