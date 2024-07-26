package models

import (
	"time"
)

// Book represents a book in the database
type Book struct {
	ID            uint       `json:"id" gorm:"primaryKey" example:"1"`
	CreatedAt     time.Time  `json:"created_at" example:"2023-04-23T18:25:43.511Z"`
	UpdatedAt     time.Time  `json:"updated_at" example:"2023-04-23T18:25:43.511Z"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" swaggertype:"string" example:"null"`
	Title         string     `json:"title" example:"The Go Programming Language"`
	Author        string     `json:"author" example:"Alan A. A. Donovan"`
	DatePublished string     `json:"date_published" example:"2015-10-26"`
	CoverImageURL string     `json:"cover_image_url" example:"https://example.com/go.jpg"`
}
