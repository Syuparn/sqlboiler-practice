package product

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/syuparn/sqlboilerpractice/domain"
	"github.com/syuparn/sqlboilerpractice/models"
	"github.com/volatiletech/null/v8"
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

func (s *summarizeProductService) Summarize(ctx context.Context, category *domain.Category) (*domain.CategoryStatistics, error) {
	// TODO: can this query be replaced with join and group by?
	n, err := models.Products(
		models.ProductWhere.CategoryID.EQ(null.StringFrom(string(category.ID))),
	).Count(ctx, s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get the number of products: %w", err)
	}

	if n < 0 {
		return nil, fmt.Errorf("the number of products must be positive but got %d", n)
	}

	return &domain.CategoryStatistics{
		CategoryID:   category.ID,
		CategoryName: category.Name,
		NumProducts:  uint(n),
	}, nil
}
