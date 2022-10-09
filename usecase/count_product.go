package usecase

import (
	"context"
	"fmt"

	"github.com/syuparn/sqlboilerpractice/domain"
)

type CountProductInputData struct{}

type CountProductOutputData struct {
	CategoryStatistics []*domain.CategoryStatistics
}

type CountProductInputPort interface {
	Handle(context.Context, *CountProductInputData) (*CountProductOutputData, error)
}

type countProductInteractor struct {
	SummarizeProductService domain.SummarizeProductService
}

func NewCountProductInputPort(
	summarizeProductService domain.SummarizeProductService,
) CountProductInputPort {
	return &countProductInteractor{
		SummarizeProductService: summarizeProductService,
	}
}

func (i *countProductInteractor) Handle(
	ctx context.Context,
	in *CountProductInputData,
) (*CountProductOutputData, error) {
	stats, err := i.SummarizeProductService.Summarize(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to summarize products: %w", err)
	}

	return &CountProductOutputData{
		CategoryStatistics: stats,
	}, nil
}
