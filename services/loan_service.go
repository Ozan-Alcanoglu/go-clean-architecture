package services

import (
	"gobackend/domain/dtos"
	"gobackend/domain/interfaces"
	"gobackend/domain/models"

	"github.com/google/uuid"
)

type loanService struct {
	repo interfaces.LoanRepository
}

func NewLoanService(repo interfaces.LoanRepository) interfaces.LoanService {

	return &loanService{repo: repo}

}

func (ls *loanService) CreateLoan(loandto *dtos.CreateLoanDto) (*dtos.LoanResponseDto, error) {
	loan := models.Loan{
		UserID:     loandto.UserID,
		BookID:     loandto.BookID,
		LoanDate:   loandto.LoanDate,
		ReturnDate: loandto.ReturnDate,
		Returned:   loandto.Returned,
	}

	if err := ls.repo.CreateLoan(&loan); err != nil {
		return nil, err
	}

	return &dtos.LoanResponseDto{
		Id:         loan.ID,
		UserID:     loan.UserID,
		BookID:     loan.BookID,
		LoanDate:   loan.LoanDate,
		ReturnDate: loan.ReturnDate,
		Returned:   loan.Returned,
	}, nil
}

func (ls *loanService) GetLoanById(id uuid.UUID) (*dtos.LoanResponseDto, error) {

	loan, err := ls.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	response := dtos.LoanResponseDto{

		Id:         loan.ID,
		UserID:     loan.UserID,
		BookID:     loan.BookID,
		LoanDate:   loan.LoanDate,
		ReturnDate: loan.ReturnDate,
		Returned:   loan.Returned,
	}

	return &response, nil

}

func (ls *loanService) UpdateLoan(id uuid.UUID, updateLoanDto *dtos.UpdateLoanDto) (*dtos.LoanResponseDto, error) {

	loan, err := ls.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	if updateLoanDto.UserID != nil {
		loan.UserID = *updateLoanDto.UserID
	}
	if updateLoanDto.BookID != nil {
		loan.BookID = *updateLoanDto.BookID
	}
	if updateLoanDto.LoanDate != nil {
		loan.LoanDate = *updateLoanDto.LoanDate
	}
	if updateLoanDto.ReturnDate != nil {
		loan.ReturnDate = *updateLoanDto.ReturnDate
	}
	if updateLoanDto.Returned != nil {
		loan.Returned = *updateLoanDto.Returned
	}

	updated, err := ls.repo.UpdateLoan(loan)

	if err != nil {
		return nil, err
	}

	response := &dtos.LoanResponseDto{
		Id:         updated.ID,
		UserID:     updated.UserID,
		BookID:     updated.BookID,
		LoanDate:   updated.LoanDate,
		ReturnDate: updated.ReturnDate,
		Returned:   updated.Returned,
	}

	return response, nil

}

func (ls *loanService) DeleteLoan(id uuid.UUID) error {
	return ls.repo.DeleteLoan(id)
}
