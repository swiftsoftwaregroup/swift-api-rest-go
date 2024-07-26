package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title         string `json:"title"`
	Author        string `json:"author"`
	DatePublished string `json:"date_published"`
	CoverImageURL string `json:"cover_image_url"`
}
