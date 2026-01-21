package interfaces

import (
	"gobackend/domain/dtos"

	"github.com/google/uuid"
)

type LoanService interface {
	CreateLoan(createLoanDto *dtos.CreateLoanDto) (*dtos.LoanResponseDto, error)
	GetLoanById(id uuid.UUID) (*dtos.LoanResponseDto, error)
	UpdateLoan(id uuid.UUID, updateLoanDto *dtos.UpdateLoanDto) (*dtos.LoanResponseDto, error)
	DeleteLoan(id uuid.UUID) error
}
