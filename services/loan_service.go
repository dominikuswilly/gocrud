package services

import (
	"context"
	"gocrud/controllers/dto"
	repo "gocrud/repository"
)

type LoanServiceInterface interface {
	Create(ctx context.Context, request dto.CreateLoanRequest) error
	// ReadById(ctx context.Context, email string) (*dto.User, error)
	// UpdateById(ctx context.Context, request dto.UpdateUserRequest, email string) error
	// DeleteById(ctx context.Context, email string) error
}

type LoanService struct {
	loanRepo repo.LoanRepositoryInterface
}

func newLoanService(loanRepo repo.LoanRepositoryInterface) LoanServiceInterface {
	return &LoanService{
		loanRepo,
	}
}

func NewLoanService(loanRepo repo.LoanRepositoryInterface) LoanServiceInterface {
	return &LoanService{
		loanRepo,
	}
}

func (s *LoanService) Create(ctx context.Context, request dto.CreateUserRequest) error {
	if err := s.loanRepo.Create(ctx, request); err != nil {
		return err
	}
	return nil
}
