package repository

import (
	"context"
	"gocrud/controllers/dto"

	"github.com/go-pg/pg/v10"
)

type LoanRepositoryInterface interface {
	Create(ctx context.Context, request dto.CreateLoanRequest) error
}

type LoanRepository struct {
	db *pg.DB
}

func NewLoanRepository(db *pg.DB) LoanRepositoryInterface {
	return &LoanRepository{
		db,
	}
}

func (r *LoanRepository) Create(ctx context.Context, request dto.CreateUserRequest) error {
	u := dto.Loan{
		request.Name,
		request.Email,
	}
	_, err := r.db.Model(&u).Context(ctx).Insert()
	if err != nil {
		return err
	}
	return nil
}
