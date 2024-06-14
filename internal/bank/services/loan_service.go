package services

import (
	"context"
	"nghiack7/bank-it/internal/bank/dtos"
)

type LoanService interface {
	CreateLoan(context.Context, *dtos.LoanRequest) (dtos.LoanResponse, int, error)
}

type loanService struct {
	// loanRepository repositories.LoanRepository
}

func NewLoanService() *loanService {
	return &loanService{}
}

func (s *loanService) CreateLoan(ctx context.Context, loan *dtos.LoanRequest) (dtos.LoanResponse, int, error) {
	loanEntity := dtos.ToEntity(loan, dtos.NewLoanEntity)
	loanEntity.Calculate()
	// need more to add calculation to database
	return dtos.ToEntity(loanEntity, dtos.CreateResponse), 200, nil
}
