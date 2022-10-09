package usecase

import (
	"context"
	"fmt"

	"github.com/syuparn/sqlboilerpractice/domain"
)

type ListCategoryInputData struct {
}

type ListCategoryOutputData struct {
	Categories []*domain.Category
}

type ListCategoryInputPort interface {
	Handle(context.Context, *ListCategoryInputData) (*ListCategoryOutputData, error)
}

type listCategoryInteractor struct {
	CategoryRepository domain.CategoryRepository
}

func NewListCategoryInputPort(
	categoryRepository domain.CategoryRepository,
) ListCategoryInputPort {
	return &listCategoryInteractor{
		CategoryRepository: categoryRepository,
	}
}

func (i *listCategoryInteractor) Handle(
	ctx context.Context,
	_ *ListCategoryInputData,
) (*ListCategoryOutputData, error) {
	categories, err := i.CategoryRepository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list category: %w", err)
	}

	return &ListCategoryOutputData{
		Categories: categories,
	}, nil
}
