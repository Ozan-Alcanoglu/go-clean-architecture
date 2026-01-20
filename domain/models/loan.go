package models

import (
	"time"

	"github.com/google/uuid"
)

type Loan struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User       User      `gorm:"foreignKey:UserID" json:"user"`
	BookID     uuid.UUID `gorm:"type:uuid" json:"book_id"`
	Book       Book      `gorm:"foreignKey:BookID" json:"book"`
	LoanDate   time.Time `gorm:"default:now()" json:"loan_date"`
	ReturnDate time.Time `gorm:"default:now()" json:"return_date"`
	Returned   bool      `json:"returned"`
}
