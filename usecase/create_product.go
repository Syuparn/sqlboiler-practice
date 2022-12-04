package usecase

import (
	"context"
	"fmt"

	"github.com/syuparn/sqlboilerpractice/domain"
)

type CreateProductInputData struct {
	Name       string
	Price      int64
	CategoryID string
}

type CreateProductOutputData struct{}

type CreateProductInputPort interface {
	Handle(context.Context, *CreateProductInputData) (*CreateProductOutputData, error)
}

type createProductInteractor struct {
	ProductRepository domain.ProductRepository
}

func NewCreateProductInputPort(
	categoryRepository domain.ProductRepository,
) CreateProductInputPort {
	return &createProductInteractor{
		ProductRepository: categoryRepository,
	}
}

func (i *createProductInteractor) Handle(
	ctx context.Context,
	in *CreateProductInputData,
) (*CreateProductOutputData, error) {
	category, err := domain.NewProduct(
		domain.ProductName(in.Name),
		domain.ProducePriceYen(in.Price),
		domain.CategoryID(in.CategoryID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	err = i.ProductRepository.Register(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("failed to register category: %w", err)
	}

	return &CreateProductOutputData{}, nil
}
