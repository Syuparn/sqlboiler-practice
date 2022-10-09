package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/syuparn/sqlboilerpractice/domain"
)

type summarizeProductService struct {
	db *sql.DB
}

var _ domain.SummarizeProductService = &summarizeProductService{}

func NewSummarizeProductService(db *sql.DB) domain.SummarizeProductService {
	return &summarizeProductService{
		db: db,
	}
}

func (s *summarizeProductService) Summarize(ctx context.Context) ([]*domain.CategoryStatistics, error) {
	return nil, errors.New("err")
}
