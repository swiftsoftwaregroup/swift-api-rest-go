package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"swift-api-rest-go/models"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Book{})
	return db
}

func TestCreateBook(t *testing.T) {
	db = setupTestDB()
	r := setupRouter()

	book := models.Book{
		Title:         "Test Book",
		Author:        "Test Author",
		DatePublished: "2023-01-01",
		CoverImageURL: "http://example.com/cover.jpg",
	}
	jsonValue, _ := json.Marshal(book)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var response models.Book
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, response.ID)
	assert.Equal(t, book.Title, response.Title)
	assert.Equal(t, book.Author, response.Author)
	assert.Equal(t, book.DatePublished, response.DatePublished)
	assert.Equal(t, book.CoverImageURL, response.CoverImageURL)
}

func TestGetBooks(t *testing.T) {
	db = setupTestDB()
	r := setupRouter()

	// Add a test book
	db.Create(&models.Book{Title: "Test Book", Author: "Test Author"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response []models.Book
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, "Test Book", response[0].Title)
	assert.Equal(t, "Test Author", response[0].Author)
}

func TestGetBook(t *testing.T) {
	db = setupTestDB()
	r := setupRouter()

	// Add a test book
	testBook := models.Book{Title: "Test Book", Author: "Test Author"}
	db.Create(&testBook)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response models.Book
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, testBook.Title, response.Title)
	assert.Equal(t, testBook.Author, response.Author)
}

func TestUpdateBook(t *testing.T) {
	db = setupTestDB()
	r := setupRouter()

	// Add a test book
	db.Create(&models.Book{Title: "Test Book", Author: "Test Author"})

	updatedBook := models.Book{
		Title:  "Updated Book",
		Author: "Updated Author",
	}
	jsonValue, _ := json.Marshal(updatedBook)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/books/1", bytes.NewBuffer(jsonValue))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response models.Book
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, updatedBook.Title, response.Title)
	assert.Equal(t, updatedBook.Author, response.Author)
}

func TestDeleteBook(t *testing.T) {
	db = setupTestDB()
	r := setupRouter()

	// Add a test book
	db.Create(&models.Book{Title: "Test Book", Author: "Test Author"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/books/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// Verify the book is deleted
	var book models.Book
	result := db.First(&book, 1)
	assert.Error(t, result.Error)
	assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
}
