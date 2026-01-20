package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Title         string    `json:"title"`
	AuthorID      uuid.UUID `gorm:"type:uuid" json:"author_id"`
	Author        Author    `gorm:"foreignKey:AuthorID" json:"author"`
	PublishedYear int       `json:"published_year"`
	Stock         int       `json:"stock"`
	CreatedAt     time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:now()" json:"updated_at"`

	Genres []Genre `gorm:"many2many:book_genres" json:"genres"`
}
