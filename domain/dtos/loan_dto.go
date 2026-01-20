package dtos

import (
	"time"

	"github.com/google/uuid"
)

type CreateLoanDto struct {
	UserID     uuid.UUID `json:"user_id"`
	BookID     uuid.UUID `json:"book_id"`
	LoanDate   time.Time `json:"loan_date"`
	ReturnDate time.Time `json:"return_date"`
	Returned   bool      `json:"returned"`
}

type UpdateLoanDto struct {
	UserID     *uuid.UUID `json:"user_id"`
	BookID     *uuid.UUID `json:"book_id"`
	LoanDate   *time.Time `json:"loan_date"`
	ReturnDate *time.Time `json:"return_date"`
	Returned   *bool      `json:"returned"`
}

type LoanResponseDto struct {
	Id         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	BookID     uuid.UUID `json:"book_id"`
	LoanDate   time.Time `json:"loan_date"`
	ReturnDate time.Time `json:"return_date"`
	Returned   bool      `json:"returned"`
}
