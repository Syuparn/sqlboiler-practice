package usecase

import (
	"context"
	"fmt"

	"github.com/syuparn/sqlboilerpractice/domain"
)

type CountProductInputData struct {
	CategoryID string
}

type CountProductOutputData struct {
	CategoryStatistics *domain.CategoryStatistics
}

type CountProductInputPort interface {
	Handle(context.Context, *CountProductInputData) (*CountProductOutputData, error)
}

type countProductInteractor struct {
	CategoryRepository      domain.CategoryRepository
	SummarizeProductService domain.SummarizeProductService
}

func NewCountProductInputPort(
	categoryRepository domain.CategoryRepository,
	summarizeProductService domain.SummarizeProductService,
) CountProductInputPort {
	return &countProductInteractor{
		CategoryRepository:      categoryRepository,
		SummarizeProductService: summarizeProductService,
	}
}

func (i *countProductInteractor) Handle(
	ctx context.Context,
	in *CountProductInputData,
) (*CountProductOutputData, error) {
	category, err := i.CategoryRepository.Get(ctx, domain.CategoryID(in.CategoryID))
	if err != nil {
		return nil, fmt.Errorf("failed to get category: %w", err)
	}

	stats, err := i.SummarizeProductService.Summarize(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("failed to summarize products: %w", err)
	}

	return &CountProductOutputData{
		CategoryStatistics: stats,
	}, nil
}
