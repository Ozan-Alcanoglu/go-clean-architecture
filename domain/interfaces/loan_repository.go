package interfaces

import (
	"gobackend/domain/models"

	"github.com/google/uuid"
)

type LoanRepository interface {
	CreateLoan(Loan *models.Loan) error
	UpdateLoan(Loan *models.Loan) (*models.Loan, error)
	DeleteLoan(id uuid.UUID) error
	GetById(id uuid.UUID) (*models.Loan, error)
}
