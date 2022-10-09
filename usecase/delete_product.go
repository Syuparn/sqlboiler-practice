package usecase

import (
	"context"
	"fmt"

	"github.com/syuparn/sqlboilerpractice/domain"
)

type DeleteProductInputData struct {
	ID string
}

type DeleteProductOutputData struct{}

type DeleteProductInputPort interface {
	Handle(context.Context, *DeleteProductInputData) (*DeleteProductOutputData, error)
}

type deleteProductInteractor struct {
	ProductRepository domain.ProductRepository
}

func NewDeleteProductInputPort(
	productRepository domain.ProductRepository,
) DeleteProductInputPort {
	return &deleteProductInteractor{
		ProductRepository: productRepository,
	}
}

func (i *deleteProductInteractor) Handle(
	ctx context.Context,
	in *DeleteProductInputData,
) (*DeleteProductOutputData, error) {
	product, err := i.ProductRepository.Get(ctx, domain.ProductID(in.ID))
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	err = i.ProductRepository.Delete(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("failed to delete product: %w", err)
	}

	return &DeleteProductOutputData{}, nil
}
