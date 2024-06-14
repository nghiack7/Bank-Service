package repositories

import (
	"context"
	"nghiack7/bank-it/internal/bank/entities"

	"gorm.io/gorm"
)

type LoanRepository interface {
	CreateLoan(context.Context, *entities.Loan) error
}

type loanRepository struct {
	db *gorm.DB
}

func (repo loanRepository) CreateLoan(ctx context.Context, loan *entities.Loan) error {
	return repo.db.Create(loan).Error
}
