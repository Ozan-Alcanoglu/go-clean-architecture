package repositories

import (
	"gobackend/domain/interfaces"
	"gobackend/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) interfaces.LoanRepository {
	return &loanRepository{db: db}
}

func (br *loanRepository) CreateLoan(Loan *models.Loan) error {

	return br.db.Create(Loan).Error

}

func (br *loanRepository) UpdateLoan(Loan *models.Loan) (*models.Loan, error) {

	if err := br.db.Save(&Loan).Error; err != nil {

		return nil, err
	}

	var updated models.Loan

	if err := br.db.Where("id = ?", Loan.ID).First(&updated).Error; err != nil {

		return nil, err

	}

	return &updated, nil
}

func (br *loanRepository) DeleteLoan(id uuid.UUID) error {

	if err := br.db.Where("id = ?", id).Delete(&models.Loan{}).Error; err != nil {
		return err
	}

	return nil
}

func (br *loanRepository) GetById(id uuid.UUID) (*models.Loan, error) {

	var loan models.Loan

	if err := br.db.Where("id = ?", id).First(&loan).Error; err != nil {
		return nil, err
	}

	return &loan, nil

}
