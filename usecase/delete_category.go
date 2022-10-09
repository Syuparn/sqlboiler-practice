package usecase

import (
	"context"
	"fmt"

	"github.com/syuparn/sqlboilerpractice/domain"
)

type DeleteCategoryInputData struct {
	ID string
}

type DeleteCategoryOutputData struct{}

type DeleteCategoryInputPort interface {
	Handle(context.Context, *DeleteCategoryInputData) (*DeleteCategoryOutputData, error)
}

type deleteCategoryInteractor struct {
	CategoryRepository domain.CategoryRepository
}

func NewDeleteCategoryInputPort(
	categoryRepository domain.CategoryRepository,
) DeleteCategoryInputPort {
	return &deleteCategoryInteractor{
		CategoryRepository: categoryRepository,
	}
}

func (i *deleteCategoryInteractor) Handle(
	ctx context.Context,
	in *DeleteCategoryInputData,
) (*DeleteCategoryOutputData, error) {
	category, err := i.CategoryRepository.Get(ctx, domain.CategoryID(in.ID))
	if err != nil {
		return nil, fmt.Errorf("failed to get category: %w", err)
	}

	err = i.CategoryRepository.Delete(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("failed to delete category: %w", err)
	}

	return &DeleteCategoryOutputData{}, nil
}
