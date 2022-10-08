package usecase

import (
	"context"
	"fmt"

	"github.com/syuparn/sqlboilerpractice/domain"
)

type CreateCategoryInputData struct {
	Name string
}

type CreateCategoryOutputData struct{}

type CreateCategoryInputPort interface {
	Handle(context.Context, *CreateCategoryInputData) (*CreateCategoryOutputData, error)
}

type createCategoryInteractor struct {
	CategoryRepository domain.CategoryRepository
}

func NewCreateCategoryInputPort(
	categoryRepository domain.CategoryRepository,
) CreateCategoryInputPort {
	return &createCategoryInteractor{
		CategoryRepository: categoryRepository,
	}
}

func (i *createCategoryInteractor) Handle(
	ctx context.Context,
	in *CreateCategoryInputData,
) (*CreateCategoryOutputData, error) {
	category, err := domain.NewCategory(domain.CategoryName(in.Name))
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	err = i.CategoryRepository.Register(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("failed to register category: %w", err)
	}

	return &CreateCategoryOutputData{}, nil
}
