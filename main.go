package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"swift-api-rest-go/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "swift-api-rest-go/swag"
)

var db *gorm.DB

// @title Book Management API
// @version 1.0
// @description A simple API for managing books.
// @host localhost:8001
// @BasePath /

func main() {
	initDB()
	r := setupRouter()

	port := ":8001"
	log.Printf("Server is starting on port%s\n", port)
	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func initDB() {
	var err error
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = ":memory:"
		log.Println("Using in-memory SQLite database")
	} else {
		log.Printf("Using SQLite database: %s\n", dbURL)
	}

	db, err = gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&models.Book{})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Add CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	r.Use(cors.New(config))

	r.POST("/books", createBook)
	r.GET("/books", getBooks)
	r.GET("/books/:id", getBook)
	r.PUT("/books/:id", updateBook)
	r.DELETE("/books/:id", deleteBook)

	// Serve the Swagger JSON
	r.GET("/swagger.json", func(c *gin.Context) {
		c.File("./swag/swagger.json")
	})

	// Swagger UI
	swaggerOpts := middleware.SwaggerUIOpts{
		SpecURL: "/swagger.json",
		Path:    "/docs",
	}
	swaggerHandler := middleware.SwaggerUI(swaggerOpts, nil)
	r.GET("/docs", gin.WrapH(swaggerHandler))

	// ReDoc
	redocOpts := middleware.RedocOpts{
		SpecURL: "/swagger.json",
		Path:    "/redoc",
	}
	redocHandler := middleware.Redoc(redocOpts, nil)
	r.GET("/redoc", gin.WrapH(redocHandler))

	return r
}

// @Summary Create a new book
// @Description Create a new book with the provided details
// @Accept json
// @Produce json
// @Param book body models.Book true "Book object"
// @Success 201 {object} models.Book
// @Failure 400 {object} map[string]string
// @Router /books [post]
func createBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&book)
	c.JSON(http.StatusCreated, book)
}

// @Summary Get all books
// @Description Get a list of all books
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func getBooks(c *gin.Context) {
	var books []models.Book
	db.Find(&books)
	c.JSON(http.StatusOK, books)
}

// @Summary Get a book by ID
// @Description Get details of a specific book by its ID
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} map[string]string
// @Router /books/{id} [get]
func getBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// @Summary Update a book
// @Description Update details of a specific book by its ID
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Updated book object"
// @Success 200 {object} models.Book
// @Failure 404 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /books/{id} [put]
func updateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&book)
	c.JSON(http.StatusOK, book)
}

// @Summary Delete a book
// @Description Delete a specific book by its ID
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /books/{id} [delete]
func deleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
